<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { ref } from 'vue'
import Header from '@/components/Header.vue'
import Banner from '@/components/Banner.vue'
import DateSelector from '@/components/DateSelector.vue'
import TimeSelector from '@/components/TimeSelector.vue'
import StaySearch from '@/components/StaySearch.vue'
import FlightSearch from '@/components/FlightSearch.vue'
import CarSearch from '@/components/CarSearch.vue'
import AttractionSearch from '@/components/AttractionSearch.vue'

const banner_bg = ref(1)

</script>

<template>
  <div class="">
    <Header />
    <Banner :option="banner_bg">
      <!-- <div><DateSelector /> <TimeSelector /></div>
      <div><DateSelector /> <TimeSelector /></div> -->
      <Tabs default-value="stays" class="">
        <TabsList class="grid w-[400px] grid-cols-4">
          <TabsTrigger value="stays" @click="banner_bg=1"> Stays </TabsTrigger>
          <TabsTrigger value="flights" @click="banner_bg=2"> Flights </TabsTrigger>
          <TabsTrigger value="cars" @click="banner_bg=3"> Cars </TabsTrigger>
          <TabsTrigger value="attractions" @click="banner_bg=4"> Attractions </TabsTrigger>
        </TabsList>
        <TabsContent value="stays"><StaySearch/></TabsContent>
        <TabsContent value="flights"><FlightSearch/></TabsContent>
        <TabsContent value="cars"><CarSearch/></TabsContent>
        <TabsContent value="attractions"><AttractionSearch/></TabsContent>
      </Tabs>
    </Banner>
  </div>
</template>
