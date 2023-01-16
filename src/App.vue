<template>
  <v-app>
    <v-main>
      <Navbar />
      <Fixtures />
    </v-main>
  </v-app>
</template>

<script>
import Navbar from './components/Navbar.vue'
import Fixtures from './components/Fixtures.vue'

export default {
  name: 'App',
  components: {
    Navbar,
    Fixtures
  }
}
</script>