<script setup lang="ts">
import {ref} from 'vue'
import MainNav from './MainNav.vue'
import Search from './Search.vue'
import UserNav from './UserNav.vue';
import Logo from './Logo.vue';

</script>

<template>
    <div class="border-b">
      <div class="flex h-16 items-center px-4">
        <Logo />
        <MainNav class="mx-6" />
        <div class="ml-auto flex items-center space-x-4">
          <Search />
          <UserNav />
        </div>
      </div>
    </div>
</template>