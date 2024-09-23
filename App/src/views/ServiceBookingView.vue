<script setup lang="ts">
import { ref } from 'vue'
import LongBookingCard from "@/components/LongBookingCard.vue"
import BookingCard from "@/components/BookingCard.vue"
</script>

<template>
    <main>
        <div v-for="(value, key) in 10" :key="key">
            <LongBookingCard />
        </div>
        
    </main>
</template>