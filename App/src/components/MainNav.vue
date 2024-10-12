<script setup lang="ts">
import { cn } from '@/lib/utils'
import { RouterLink } from 'vue-router'

const navs = [
  {
    name: 'Overview',
    href: '/examples/dashboard',
    class: 'text-sm font-medium transition-colors hover:text-primary'
  },
  {
    name: 'Customers',
    href: '/examples/dashboard',
    class: 'text-sm font-medium text-muted-foreground transition-colors hover:text-primary'
  },
  {
    name: 'Products',
    href: '/examples/dashboard',
    class: 'text-sm font-medium text-muted-foreground transition-colors hover:text-primary'
  },
  {
    name: 'Settings',
    href: '/examples/dashboard',
    class: 'text-sm font-medium text-muted-foreground transition-colors hover:text-primary'
  }
]
</script>

<template>
  <nav :class="cn('flex items-center space-x-4 lg:space-x-6', $attrs.class ?? '')">
    <RouterLink v-for="(item, key) in navs" :key="key" :class="item.class" :to="item.href">
      {{ item.name }}
    </RouterLink>
  </nav>
</template>
