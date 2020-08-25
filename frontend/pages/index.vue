<template>
  <v-container>
    <v-row justify="center" justify-md="start">
      <v-col cols="12" md="3">
        <dashboard-item
          ref="heartRateChart"
          title="Heart Rate"
          :value="heartRate"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import DashboardItem from '@/components/DashboardItem.vue'
import { APIClient } from '@/api/api_client'
import { GetHeartRateRequest } from '@/api/classes/GetHeartRateRequest'

@Component({
  components: { DashboardItem },
})
export default class PageIndex extends Vue {
  heartRate: number = 0
  apiClient!: APIClient

  get refs(): any {
    return this.$refs
  }

  fetchData() {
    this.apiClient.getHeartRate(new GetHeartRateRequest()).then((res) => {
      if (res.status) {
        if (Number(res.message) !== this.heartRate) {
          this.heartRate = Number(res.message)
        } else {
          this.refs.heartRateChart.onChangeValue()
        }
      }
    })
  }

  mounted() {
    this.apiClient = new APIClient('', {}, process.env.BASE_URL, {})

    this.fetchData()
    window.setInterval(this.fetchData, 5000)
  }

  head() {
    return {
      title: 'Home',
    }
  }
}
</script>
