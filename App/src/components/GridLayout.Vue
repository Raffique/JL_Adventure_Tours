<script setup lang="ts">
import { ref, computed } from 'vue'

const props = withDefaults(
  defineProps<{
    lgCols?: number,
    mdCols?: number,
  }>(),
  {
    lgCols: 4,
    mdCols: 2
  }
)

</script>

<template>
    <div class="grid gap-6 px-4" 
        :class="`lg:grid-cols-${props.lgCols} md:grid-cols-${props.mdCols}`" 
    >
        <slot />
    </div>
</template>
