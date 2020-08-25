<template>
  <v-card>
    <dashboard-chart :chart-data="chartData" :chart-options="chartOptions" />
    <v-card-title>{{ title }}</v-card-title>
    <v-card-text>
      <p class="">{{ value }}</p>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from 'nuxt-property-decorator'
import { ChartData, ChartOptions, ChartPoint } from 'chart.js'
import DashboardChart from '~/components/DashboardChart.vue'
import 'chartjs-plugin-streaming'

@Component({
  components: { DashboardChart },
})
export default class DashboardItem extends Vue {
  @Prop()
  title!: string

  @Prop()
  value!: number

  chartData: ChartData = {
    datasets: [
      {
        label: this.title,
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
        borderColor: 'rgb(255, 99, 132)',
        data: [{ t: new Date(), y: this.value }],
      },
    ],
  }

  chartOptions: ChartOptions = {
    scales: {
      xAxes: [
        {
          type: 'realtime',
          ticks: {
            stepSize: 10,
          },
        },
      ],
      yAxes: [
        {
          ticks: {
            min: 60,
            max: 160,
          },
        },
      ],
    },
    plugins: {
      streaming: {
        duration: 20000,
        refresh: 2500,
        delay: 6000,
      },
    },
  }

  @Watch('value')
  onChangeValue() {
    if (
      this.chartData.datasets !== undefined &&
      this.chartData.datasets[0].data !== undefined
    ) {
      const p: ChartPoint = { t: new Date(), y: this.value }
      ;(this.chartData.datasets[0].data as ChartPoint[]).push(p)
    }
  }
}
</script>
