<template>
  <div class="flex justify-between pt-5 px-7">
    <button
      class="px-1 py-1 inline-flex text-xl items-center"
      type="button"
      @click="toggle"
    >
      {{ selected.name }}
      <div class="pi pi-chevron-down text-sm ml-1.5" />
    </button>
    <Menu
      ref="menu"
      :model="items"
      :popup="true"
    />
  </div>
</template>

<script setup lang="ts">
import Menu from "primevue/menu"
import { ref } from "vue"
import { useRouter } from "vue-router"
import { topNavigationItems } from "./routes"

const menuItems = topNavigationItems.map(({ path, name }) => ({
    label: name,
    url: path,
}))
const items = ref(menuItems)
const menu = ref<Menu | null>(null)
const router = useRouter()
const currentPath = router.currentRoute.value.path
const selected = topNavigationItems.find(({ key }) => {
  return key ? currentPath.startsWith(key) : false
}) ?? topNavigationItems[0]

function toggle(event: PointerEvent) {
    menu.value?.toggle(event)
}
</script>