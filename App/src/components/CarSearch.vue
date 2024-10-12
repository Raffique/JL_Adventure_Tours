<script setup lang="ts">
import { ref } from 'vue'
import { Input } from '@/components/ui/input'
import Button from '@/components/ui/button/Button.vue'
import DateSelector from '@/components/DateSelector.vue'
import TimeSelector from '@/components/TimeSelector.vue'
</script>

<template>
  <div class="flex">
    <div class="flex">
      <Input />
      <DateSelector class="inline-flex" />
      <TimeSelector class="inline-flex" />
    </div>
    <div class="flex">
      <Input />
      <DateSelector class="inline-flex" />
      <TimeSelector class="inline-flex" />
    </div>
    <Button class="bg-blue-700">Search</Button>
  </div>
</template>
