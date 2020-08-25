<script lang="ts">
import { Component, Prop } from 'nuxt-property-decorator'
import { Mixins } from 'vue-property-decorator'
import { Line, mixins } from 'vue-chartjs'
import Chart from 'chart.js'

@Component
export default class DashboardChart extends Mixins(Line, mixins.reactiveProp) {
  @Prop()
  chartData!: Chart.ChartData

  @Prop()
  chartOptions!: Chart.ChartOptions

  mounted() {
    this.renderChart(this.chartData, this.chartOptions)
  }
}
</script>
