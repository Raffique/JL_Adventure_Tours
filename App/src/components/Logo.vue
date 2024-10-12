<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink } from 'vue-router';
</script>

<template>
  <div>
    <RouterLink to="/" class="-m-1.5 p-1.5">
      <span class="sr-only">Your Company</span>
      <img
        class="h-8 w-auto"
        src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
        alt=""
      />
    </RouterLink>
  </div>
</template>
