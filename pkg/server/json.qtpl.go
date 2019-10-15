// Code generated by qtc from "json.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pkg/server/json.qtpl:2
package server

//line pkg/server/json.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pkg/server/json.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pkg/server/json.qtpl:2
func streamhttpError(qw422016 *qt422016.Writer, error *HttpError) {
//line pkg/server/json.qtpl:2
	qw422016.N().S(`{"error":`)
//line pkg/server/json.qtpl:4
	qw422016.N().Q(error.Message)
//line pkg/server/json.qtpl:4
	qw422016.N().S(`}`)
//line pkg/server/json.qtpl:6
}

//line pkg/server/json.qtpl:6
func writehttpError(qq422016 qtio422016.Writer, error *HttpError) {
//line pkg/server/json.qtpl:6
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:6
	streamhttpError(qw422016, error)
//line pkg/server/json.qtpl:6
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:6
}

//line pkg/server/json.qtpl:6
func httpError(error *HttpError) string {
//line pkg/server/json.qtpl:6
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:6
	writehttpError(qb422016, error)
//line pkg/server/json.qtpl:6
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:6
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:6
	return qs422016
//line pkg/server/json.qtpl:6
}

//line pkg/server/json.qtpl:8
func StreamGroupedMetricList(qw422016 *qt422016.Writer, list []MedianResult) {
//line pkg/server/json.qtpl:8
	qw422016.N().S(`{"groupNames": [`)
//line pkg/server/json.qtpl:11
	if len(list) > 0 {
//line pkg/server/json.qtpl:12
		for i, value := range list[0].groupedValues {
//line pkg/server/json.qtpl:13
			if i != 0 {
//line pkg/server/json.qtpl:13
				qw422016.N().S(`,`)
//line pkg/server/json.qtpl:13
			}
//line pkg/server/json.qtpl:13
			qw422016.N().S(`"`)
//line pkg/server/json.qtpl:14
			qw422016.N().S(value.group)
//line pkg/server/json.qtpl:14
			qw422016.N().S(`"`)
//line pkg/server/json.qtpl:15
		}
//line pkg/server/json.qtpl:16
	}
//line pkg/server/json.qtpl:16
	qw422016.N().S(`],"data": [`)
//line pkg/server/json.qtpl:19
	for i, item := range list {
//line pkg/server/json.qtpl:20
		if i != 0 {
//line pkg/server/json.qtpl:20
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:20
		}
//line pkg/server/json.qtpl:20
		qw422016.N().S(`{`)
//line pkg/server/json.qtpl:22
		for _, value := range item.groupedValues {
//line pkg/server/json.qtpl:22
			qw422016.N().S(`"`)
//line pkg/server/json.qtpl:23
			qw422016.N().S(value.group)
//line pkg/server/json.qtpl:23
			qw422016.N().S(`":`)
//line pkg/server/json.qtpl:23
			qw422016.N().D(value.value)
//line pkg/server/json.qtpl:23
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:24
		}
//line pkg/server/json.qtpl:27
		qw422016.N().S(`"metric":`)
//line pkg/server/json.qtpl:28
		qw422016.N().Q(item.metricName)
//line pkg/server/json.qtpl:28
		qw422016.N().S(`}`)
//line pkg/server/json.qtpl:30
	}
//line pkg/server/json.qtpl:30
	qw422016.N().S(`]}`)
//line pkg/server/json.qtpl:33
}

//line pkg/server/json.qtpl:33
func WriteGroupedMetricList(qq422016 qtio422016.Writer, list []MedianResult) {
//line pkg/server/json.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:33
	StreamGroupedMetricList(qw422016, list)
//line pkg/server/json.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:33
}

//line pkg/server/json.qtpl:33
func GroupedMetricList(list []MedianResult) string {
//line pkg/server/json.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:33
	WriteGroupedMetricList(qb422016, list)
//line pkg/server/json.qtpl:33
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:33
	return qs422016
//line pkg/server/json.qtpl:33
}

//line pkg/server/json.qtpl:35
func StreamInfo(qw422016 *qt422016.Writer, productNames []string, essentialMetricNames []string, instantMetricNames []string, productNameToMachineNames map[string][]string) {
//line pkg/server/json.qtpl:35
	qw422016.N().S(`{"productNames":`)
//line pkg/server/json.qtpl:37
	streamsafeStringList(qw422016, productNames)
//line pkg/server/json.qtpl:37
	qw422016.N().S(`,"durationMetricNames": [`)
//line pkg/server/json.qtpl:39
	for i, name := range essentialMetricNames {
//line pkg/server/json.qtpl:40
		if i != 0 {
//line pkg/server/json.qtpl:40
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:40
		}
//line pkg/server/json.qtpl:40
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:41
		qw422016.N().S(name)
//line pkg/server/json.qtpl:41
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:42
	}
//line pkg/server/json.qtpl:42
	qw422016.N().S(`],"instantMetricNames": [`)
//line pkg/server/json.qtpl:45
	for i, name := range instantMetricNames {
//line pkg/server/json.qtpl:46
		if i != 0 {
//line pkg/server/json.qtpl:46
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:46
		}
//line pkg/server/json.qtpl:46
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:47
		qw422016.N().S(name)
//line pkg/server/json.qtpl:47
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:48
	}
//line pkg/server/json.qtpl:48
	qw422016.N().S(`],"productToMachine": {`)
//line pkg/server/json.qtpl:51
	for i, product := range productNames {
//line pkg/server/json.qtpl:52
		if i != 0 {
//line pkg/server/json.qtpl:52
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:52
		}
//line pkg/server/json.qtpl:52
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:53
		qw422016.N().S(product)
//line pkg/server/json.qtpl:53
		qw422016.N().S(`":`)
//line pkg/server/json.qtpl:53
		streamwriteMachineInfoList(qw422016, productNameToMachineNames[product])
//line pkg/server/json.qtpl:54
	}
//line pkg/server/json.qtpl:54
	qw422016.N().S(`}}`)
//line pkg/server/json.qtpl:57
}

//line pkg/server/json.qtpl:57
func WriteInfo(qq422016 qtio422016.Writer, productNames []string, essentialMetricNames []string, instantMetricNames []string, productNameToMachineNames map[string][]string) {
//line pkg/server/json.qtpl:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:57
	StreamInfo(qw422016, productNames, essentialMetricNames, instantMetricNames, productNameToMachineNames)
//line pkg/server/json.qtpl:57
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:57
}

//line pkg/server/json.qtpl:57
func Info(productNames []string, essentialMetricNames []string, instantMetricNames []string, productNameToMachineNames map[string][]string) string {
//line pkg/server/json.qtpl:57
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:57
	WriteInfo(qb422016, productNames, essentialMetricNames, instantMetricNames, productNameToMachineNames)
//line pkg/server/json.qtpl:57
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:57
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:57
	return qs422016
//line pkg/server/json.qtpl:57
}

//line pkg/server/json.qtpl:59
func streamsafeStringList(qw422016 *qt422016.Writer, list []string) {
//line pkg/server/json.qtpl:59
	qw422016.N().S(`[`)
//line pkg/server/json.qtpl:61
	for i, v := range list {
//line pkg/server/json.qtpl:62
		if i != 0 {
//line pkg/server/json.qtpl:62
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:62
		}
//line pkg/server/json.qtpl:62
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:63
		qw422016.N().S(v)
//line pkg/server/json.qtpl:63
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:64
	}
//line pkg/server/json.qtpl:64
	qw422016.N().S(`]`)
//line pkg/server/json.qtpl:66
}

//line pkg/server/json.qtpl:66
func writesafeStringList(qq422016 qtio422016.Writer, list []string) {
//line pkg/server/json.qtpl:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:66
	streamsafeStringList(qw422016, list)
//line pkg/server/json.qtpl:66
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:66
}

//line pkg/server/json.qtpl:66
func safeStringList(list []string) string {
//line pkg/server/json.qtpl:66
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:66
	writesafeStringList(qb422016, list)
//line pkg/server/json.qtpl:66
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:66
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:66
	return qs422016
//line pkg/server/json.qtpl:66
}

//line pkg/server/json.qtpl:68
func streamwriteMachineInfoList(qw422016 *qt422016.Writer, machines []string) {
//line pkg/server/json.qtpl:68
	qw422016.N().S(`[`)
//line pkg/server/json.qtpl:70
	for i, machine := range machines {
//line pkg/server/json.qtpl:71
		if i != 0 {
//line pkg/server/json.qtpl:71
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:71
		}
//line pkg/server/json.qtpl:72
		qw422016.N().Q(machine)
//line pkg/server/json.qtpl:73
	}
//line pkg/server/json.qtpl:73
	qw422016.N().S(`]`)
//line pkg/server/json.qtpl:75
}

//line pkg/server/json.qtpl:75
func writewriteMachineInfoList(qq422016 qtio422016.Writer, machines []string) {
//line pkg/server/json.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:75
	streamwriteMachineInfoList(qw422016, machines)
//line pkg/server/json.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:75
}

//line pkg/server/json.qtpl:75
func writeMachineInfoList(machines []string) string {
//line pkg/server/json.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:75
	writewriteMachineInfoList(qb422016, machines)
//line pkg/server/json.qtpl:75
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:75
	return qs422016
//line pkg/server/json.qtpl:75
}
