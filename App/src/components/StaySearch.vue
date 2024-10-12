<script setup lang="ts">
import { ref } from 'vue'
import { Input } from '@/components/ui/input'
import Button from '@/components/ui/button/Button.vue'
import DateRangeSelector from '@/components/DateRangeSelector.vue'
import TimeSelector from '@/components/TimeSelector.vue'
</script>

<template>
    <div class="flex">
        <Input/>
        <DateRangeSelector/>
        <!-- possible selector for amount of adults and amount of children -->
        <Button class="bg-blue-700">Search</Button>
    </div>
</template>