import { deepEqual } from "fast-equals"
import { ToastSeverity } from "primevue/api"
import ToastEventBus from "primevue/toasteventbus"
import { catchError, delay, distinctUntilChanged, EMPTY, mergeMap, Observable, of, retry, takeUntil, timer } from "rxjs"
import { fromFetch } from "rxjs/fetch"
import { Ref, watch } from "vue"

export function refToObservable<T>(ref: Ref<T>, deep: boolean = false): Observable<T> {
  return new Observable<T>(context => {
    watch(ref, value => {
      return context.next(value)
    }, {deep})
    context.next(ref.value)
  }).pipe(
    deep ? distinctUntilChanged(deepEqual) : distinctUntilChanged(), 
  )
}

const serverNotAvailableMessage = {
  severity: ToastSeverity.ERROR,
  summary: "Server is not available",
  detail: "Please check that server is running.",
  life: 3_000,
}

function getRequest(url: string) {
  const searchString = "/!("
  if (url.length > 2_000 && url.indexOf(searchString) > 0) {
    return new Request(url.substring(0, url.indexOf(searchString) + 1), {
      method: "POST",
      body: url.substring(url.indexOf(searchString) + 1),
    })
  }
  else {
    return new Request(url, {
      method: "GET",
    })
  }
}

export function fromFetchWithRetryAndErrorHandling<T>(url: string, unavailableErrorMessage: ({summary: string; detail: string}) | null = null): Observable<T> {
  const request = getRequest(url)
  return fromFetch(request)
    .pipe(
      // promise to result
      mergeMap(response => {
        if (response.ok) {
          return response.json() as Promise<T>
        }
        else {
          throw new Error(`cannot load (status=${response.status})`)
        }
      }),
      // timeout
      takeUntil(timer(8_000)),
      // retry at least three times
      retry({
        count: 3,
        delay(error, retryCount) {
          return of(error).pipe(delay(1_000 * retryCount))
        },
      }),
      catchError((error, _caught) => {
        if (error instanceof TypeError) {
          if (unavailableErrorMessage == null) {
            ToastEventBus.emit("remove", serverNotAvailableMessage)
            ToastEventBus.emit("add", serverNotAvailableMessage)
          }
          else {
            ToastEventBus.emit("add", {
              severity: ToastSeverity.ERROR,
              summary: unavailableErrorMessage.summary,
              detail: unavailableErrorMessage.detail,
              life: 3_000,
            })
          }
        }
        else {
          ToastEventBus.emit("add", {
            severity: ToastSeverity.ERROR,
            summary: "Cannot load",
            detail: `${(error as Error).message}`,
            life: 3_000
          })
        }
        console.error("cannot load", url, error)
        return EMPTY
      })
    )
}
