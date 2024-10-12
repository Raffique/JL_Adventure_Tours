<script setup lang="ts">
import { ref} from 'vue'

</script>

<template>
    <div id="date-range-picker" date-rangepicker class="flex items-center mt-5">
    <div class="relative">
      <input
        id="datepicker-range-start"
        name="start"
        type="date"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="Select date start"
        min="2024-09-23"
      />
      <input
        id="datepicker-range-start"
        type="time"
      />
    </div>
    <span class="mx-4 text-gray-500">to</span>
    <div class="relative">
      <input
        id="datepicker-range-end"
        name="end"
        type="date"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="Select date end"
        min="2024-09-23"
        max="2024-12-30"
      />
      <input
        id="datepicker-range-start"
        name="start"
        type="time"
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="Select date start"
        min="2024-09-23"
      />
    </div>
  </div>
</template>