<template>
  <div class="flex flex-col gap-5">
    <Toolbar class="customToolbar">
      <template #start>
        <TimeRangeSelect
          :ranges="TimeRangeConfigurator.timeRanges"
          :value="timeRangeConfigurator.value.value"
          :on-change="onChangeRange"
        >
          <template #icon>
            <CalendarIcon class="w-4 h-4 text-gray-500" />
          </template>
        </TimeRangeSelect>
        <BranchSelect
          :branch-configurator="branchConfigurator"
          :triggered-by-configurator="triggeredByConfigurator"
        />
        <DimensionHierarchicalSelect
          label="Machine"
          :dimension="machineConfigurator"
        >
          <template #icon>
            <ComputerDesktopIcon class="w-4 h-4 text-gray-500" />
          </template>
        </DimensionHierarchicalSelect>
      </template>
    </Toolbar>

    <main class="flex">
      <div
        ref="container"
        class="flex flex-1 flex-col gap-6 overflow-hidden"
      >
        <section class="flex gap-6">
          <div class="flex-1">
            <AggregationChart
              :configurators="averagesConfigurators"
              :aggregated-measure="'highlighting\_%'"
              :aggregated-project="'%\_k1'"
              :is-like="true"
              :title="'mean all project highlighting K1'"
            />
          </div>
          <div class="flex-1">
            <AggregationChart
              :configurators="averagesConfigurators"
              :aggregated-measure="'highlighting\_%'"
              :aggregated-project="'%\_k2'"
              :is-like="true"
              :title="'mean all project highlighting K2'"
            />
          </div>
        </section>
        <section class="flex gap-x-6">
          <div class="flex-1">
            <GroupProjectsChart
              label="Intellij highlight mean value with Library cache K1"
              measure="highlighting#mean_value"
              :projects="[
                'intellij_commit/highlight/KtOCSwiftChangeSignatureTest_with_library_cache_k1',
                'intellij_commit/highlight/KotlinHighLevelFunctionParameterInfoHandler_with_library_cache_k1',
                'intellij_commit/highlight/ContentManagerImpl_with_library_cache_k1',
                'intellij_commit/highlight/JdkList_with_library_cache_k1',
                'intellij_commit/highlight/ComposeCompletionContributorTest_with_library_cache_k1',
                'intellij_commit/highlight/AgpUpgradeRefactoringProcessor_with_library_cache_k1',
                'intellij_commit/highlight/AndroidModelTest_with_library_cache_k1',
                'intellij_commit/highlight/SecureWireOverStreamTransport_with_library_cache_k1',
                'intellij_commit/highlight/DexInlineCallStackComparisonTest_with_library_cache_k1',
                'intellij_commit/highlight/DexLocalVariableTableBreakpointTest_with_library_cache_k1',
                'intellij_commit/highlight/OraIntrospector_with_library_cache_k1',
                'intellij_commit/highlight/SolutionModel.Generated_with_library_cache_k1',
                'intellij_commit/highlight/UIAutomationInteractionModel.Generated_with_library_cache_k1',
              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
          <div class="flex-1">
            <GroupProjectsChart
              label="Intellij highlight mean value with Library cache K2"
              measure="highlighting#mean_value"
              :projects="[
                'intellij_commit/highlight/KtOCSwiftChangeSignatureTest_with_library_cache_k2',
                'intellij_commit/highlight/KotlinHighLevelFunctionParameterInfoHandler_with_library_cache_k2',
                'intellij_commit/highlight/ContentManagerImpl_with_library_cache_k2',
                'intellij_commit/highlight/JdkList_with_library_cache_k2',
                'intellij_commit/highlight/ComposeCompletionContributorTest_with_library_cache_k2',
                'intellij_commit/highlight/AgpUpgradeRefactoringProcessor_with_library_cache_k2',
                'intellij_commit/highlight/AndroidModelTest_with_library_cache_k2',
                'intellij_commit/highlight/SecureWireOverStreamTransport_with_library_cache_k2',
                'intellij_commit/highlight/DexInlineCallStackComparisonTest_with_library_cache_k2',
                'intellij_commit/highlight/DexLocalVariableTableBreakpointTest_with_library_cache_k2',
                'intellij_commit/highlight/OraIntrospector_with_library_cache_k2',
                'intellij_commit/highlight/SolutionModel.Generated_with_library_cache_k2',
                'intellij_commit/highlight/UIAutomationInteractionModel.Generated_with_library_cache_k2',
              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
        </section>

        <section class="flex gap-x-6">
          <div class="flex-1">
            <GroupProjectsChart
              label="'Intellij suite 2' highlight mean value with Library cache K1"
              measure="highlighting#mean_value"
              :projects="[
                'intellij_commit/highlight/DexInlineCallStackComparisonTest_with_library_cache_k1',
                'intellij_commit/highlight/DexLocalVariableTableBreakpointTest_with_library_cache_k1',
                'intellij_commit/highlight/OraIntrospector_with_library_cache_k1',
                'intellij_commit/highlight/SolutionModel.Generated_with_library_cache_k1',
                'intellij_commit/highlight/UIAutomationInteractionModel.Generated_with_library_cache_k1',
              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
          <div class="flex-1">
            <GroupProjectsChart
              label="'Intellij suite 2' highlight mean value with Library cache K2"
              measure="highlighting#mean_value"
              :projects="[
                'intellij_commit/highlight/DexInlineCallStackComparisonTest_with_library_cache_k2',
                'intellij_commit/highlight/DexLocalVariableTableBreakpointTest_with_library_cache_k2',
                'intellij_commit/highlight/OraIntrospector_with_library_cache_k2',
                'intellij_commit/highlight/SolutionModel.Generated_with_library_cache_k2',
                'intellij_commit/highlight/UIAutomationInteractionModel.Generated_with_library_cache_k2',
              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
        </section>

        <section class="flex gap-x-6">
          <div class="flex-1">
            <GroupProjectsChart
              label="'Kotlin lang' highlight mean value with Library cache K1"
              measure="highlighting#mean_value"
              :projects="[
                'kotlin_lang/highlight/CommonParser_with_library_cache_k1',
                'kotlin_lang/highlight/FirErrors_with_library_cache_k1',
                'kotlin_lang/highlight/Flag_with_library_cache_k1',
                'kotlin_lang/highlight/KtFirDataClassConverters_with_library_cache_k1',
                'kotlin_lang/highlight/DefaultArgumentStubGenerator_with_library_cache_k1',
              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
          <div class="flex-1">
            <GroupProjectsChart
              label="'Kotlin lang' highlight mean value with Library cache K2"
              measure="highlighting#mean_value"
              :projects="[
                'kotlin_lang/highlight/CommonParser_with_library_cache_k2',
                'kotlin_lang/highlight/FirErrors_with_library_cache_k2',
                'kotlin_lang/highlight/Flag_with_library_cache_k2',
                'kotlin_lang/highlight/KtFirDataClassConverters_with_library_cache_k2',
                'kotlin_lang/highlight/DefaultArgumentStubGenerator_with_library_cache_k2',
              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
        </section>

        <section class="flex gap-x-6">
          <div class="flex-1">
            <GroupProjectsChart
              label="'Kotlin language server' highlight mean value with Library cache K1"
              measure="highlighting#mean_value"
              :projects="[
                'kotlin_language_server/highlight/Compiler_with_library_cache_k1',
                'kotlin_language_server/highlight/Completions_with_library_cache_k1',
                'kotlin_language_server/highlight/CompletionsTest_with_library_cache_k1',
                'kotlin_language_server/highlight/JavaElementConverter_with_library_cache_k1',
                'kotlin_language_server/highlight/KotlinTextDocumentService_with_library_cache_k1',
                'kotlin_language_server/highlight/QuickFixesTest_with_library_cache_k1',
                'kotlin_language_server/highlight/SourcePath_with_library_cache_k1'
              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
          <div class="flex-1">
            <GroupProjectsChart
              label="'Kotlin language server' highlight mean value with Library cache K2"
              measure="highlighting#mean_value"
              :projects="[
                'kotlin_language_server/highlight/Compiler_with_library_cache_k2',
                'kotlin_language_server/highlight/Completions_with_library_cache_k2',
                'kotlin_language_server/highlight/CompletionsTest_with_library_cache_k2',
                'kotlin_language_server/highlight/JavaElementConverter_with_library_cache_k2',
                'kotlin_language_server/highlight/KotlinTextDocumentService_with_library_cache_k2',
                'kotlin_language_server/highlight/QuickFixesTest_with_library_cache_k2',
                'kotlin_language_server/highlight/SourcePath_with_library_cache_k2',


              ]"
              :server-configurator="serverConfigurator"
              :configurators="dashboardConfigurators"
            />
          </div>
        </section>
      </div>
      <InfoSidebar />
    </main>
  </div>
</template>

<script setup lang="ts">
import { PersistentStateManager } from "shared/src/PersistentStateManager"
import DimensionHierarchicalSelect from "shared/src/components/DimensionHierarchicalSelect.vue"
import { createBranchConfigurator } from "shared/src/configurators/BranchConfigurator"
import { MachineConfigurator } from "shared/src/configurators/MachineConfigurator"
import { privateBuildConfigurator } from "shared/src/configurators/PrivateBuildConfigurator"
import { ServerConfigurator } from "shared/src/configurators/ServerConfigurator"
import { TimeRangeConfigurator } from "shared/src/configurators/TimeRangeConfigurator"
import { provideReportUrlProvider } from "shared/src/lineChartTooltipLinkProvider"
import { provide, ref } from "vue"
import { useRouter } from "vue-router"
import { containerKey, sidebarVmKey } from "../../../shared/keys"
import InfoSidebar from "../../InfoSidebar.vue"
import { InfoSidebarVmImpl } from "../../InfoSidebarVm"
import AggregationChart from "../../charts/AggregationChart.vue"
import GroupProjectsChart from "../../charts/GroupProjectsChart.vue"
import BranchSelect from "../../common/BranchSelect.vue"
import TimeRangeSelect from "../../common/TimeRangeSelect.vue"

provideReportUrlProvider(false)

const dbName = "perfintDev"
const dbTable = "kotlin"
const initialMachine = "linux-blade-hetzner"
const container = ref<HTMLElement>()
const router = useRouter()
const sidebarVm = new InfoSidebarVmImpl()

provide(containerKey, container)
provide(sidebarVmKey, sidebarVm)

const serverConfigurator = new ServerConfigurator(dbName, dbTable)
const persistenceForDashboard = new PersistentStateManager("kotlinDev_dashboard", {
  machine: initialMachine,
  project: [],
  branch: "master",
}, router)

const timeRangeConfigurator = new TimeRangeConfigurator(persistenceForDashboard)

const branchConfigurator = createBranchConfigurator(serverConfigurator, persistenceForDashboard, [timeRangeConfigurator])
const machineConfigurator = new MachineConfigurator(
  serverConfigurator,
  persistenceForDashboard,
  [timeRangeConfigurator, branchConfigurator],
)
const triggeredByConfigurator = privateBuildConfigurator(
  serverConfigurator,
  persistenceForDashboard,
  [branchConfigurator, timeRangeConfigurator],
)

const dashboardConfigurators = [
  serverConfigurator,
  branchConfigurator,
  machineConfigurator,
  timeRangeConfigurator,
  triggeredByConfigurator,
]

const averagesConfigurators = [
  serverConfigurator,
  branchConfigurator,
  machineConfigurator,
  timeRangeConfigurator,
]

function onChangeRange(value: string) {
  timeRangeConfigurator.value.value = value
}
</script>

<style>
.customToolbar {
  background-color: transparent;
  border: none;
  padding: 0;
}
</style>