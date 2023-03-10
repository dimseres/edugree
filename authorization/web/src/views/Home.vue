<template>
  <DashboardTemplate>
    <RouterView/>
  </DashboardTemplate>
</template>

<script lang="ts">
import AuthTemplate from "../components/templates/AuthTemplate.vue";
import DashboardTemplate from "../components/templates/DashboardTemplate.vue";

export default {
  name: "Home.vue",
  components: {DashboardTemplate, AuthTemplate}
}
</script>