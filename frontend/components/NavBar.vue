<template>
  <header>
    <v-app-bar app dense>
      <v-app-bar-nav-icon @click="isDrawerEnabled = !isDrawerEnabled" />
      <v-toolbar-title>Dashboard</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="isDrawerEnabled" app>
      <v-list dense nav>
        <v-list-item></v-list-item>
      </v-list>
    </v-navigation-drawer>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class NavBar extends Vue {
  isDrawerEnabled: boolean = false
}
</script>
