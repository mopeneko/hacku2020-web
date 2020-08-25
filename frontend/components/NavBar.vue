<template>
  <header>
    <v-app-bar app dense>
      <v-app-bar-nav-icon @click="isDrawerEnabled = !isDrawerEnabled" />
      <v-toolbar-title>Dashboard</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="isDrawerEnabled" app>
      <v-list dense nav>
        <v-list-item v-for="item in items" :key="item.title" link :to="item.to">
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

interface Item {
  title: string
  icon: string
  to: string
}

@Component
export default class NavBar extends Vue {
  isDrawerEnabled: boolean = false
  items: Item[] = [{ title: 'Home', icon: 'mdi-home', to: '/' }]
}
</script>
