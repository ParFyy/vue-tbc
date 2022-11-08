<template>
  <Navbar />
  <Fixtures />
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