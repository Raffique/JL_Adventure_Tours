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
import GridLayout from '@/components/GridLayout.vue'

const banner_bg = ref(1)

const s1Cards = [
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
  {
    title: 'Title 1',
    description: 'lorem itusopa fghj',
    content: 'dgdffdgfdvd',
    footer: 'dfdfdfdf'
  },
]
</script>

<template>
  <div class="">
    <Header />
    <Banner :option="banner_bg">
      <!-- <div><DateSelector /> <TimeSelector /></div>
      <div><DateSelector /> <TimeSelector /></div> -->
      <Tabs default-value="stays" class="">
        <TabsList class="grid w-[400px] grid-cols-4">
          <TabsTrigger value="stays" @click="banner_bg = 1"> Stays </TabsTrigger>
          <TabsTrigger value="flights" @click="banner_bg = 2"> Flights </TabsTrigger>
          <TabsTrigger value="cars" @click="banner_bg = 3"> Cars </TabsTrigger>
          <TabsTrigger value="attractions" @click="banner_bg = 4"> Attractions </TabsTrigger>
        </TabsList>
        <TabsContent value="stays"><StaySearch /></TabsContent>
        <TabsContent value="flights"><FlightSearch /></TabsContent>
        <TabsContent value="cars"><CarSearch /></TabsContent>
        <TabsContent value="attractions"><AttractionSearch /></TabsContent>
      </Tabs>
    </Banner>

    <!-- Section 1 -->
    <h1 class="m-3 my-5 text-5xl text-blue-700 font-sans">Segment 1</h1>
    <GridLayout>
      <div v-for="(card, idx) in s1Cards.slice(0,4)" :key="idx">
        <Card class="rounded-t-3xl rounded-b-3xl bg-black">
          <CardHeader class="rounded-t-3xl bg-gray-100">
            <CardTitle>{{ card.title }}</CardTitle>
            <CardDescription>{{ card.description }}</CardDescription>
          </CardHeader>
          <CardContent class="bg-gray-100"> {{ card.content }} </CardContent>
          <CardFooter class="rounded-b-3xl bg-gray-100"> {{ card.footer }} </CardFooter>
        </Card>
      </div>
    </GridLayout>

    <!-- Section 2 -->
    <h1 class="m-3 my-5 text-5xl text-blue-700 font-sans">Segment 2</h1>
    <GridLayout :lg-cols="3" :md-cols="3" >
      <div v-for="(card, idx) in s1Cards.slice(0,6)" :key="idx">
        <Card class="rounded-t-3xl rounded-b-3xl bg-black">
          <CardHeader class="rounded-t-3xl bg-gray-100">
            <CardTitle>{{ card.title }}</CardTitle>
            <CardDescription>{{ card.description }}</CardDescription>
          </CardHeader>
          <CardContent class="bg-gray-100"> {{ card.content }} </CardContent>
          <CardFooter class="rounded-b-3xl bg-gray-100"> {{ card.footer }} </CardFooter>
        </Card>
      </div>
    </GridLayout>

    <!-- Section 3 -->
    <h1 class="m-3 my-5 text-5xl text-blue-700 font-sans">Segment 3</h1>
    <GridLayout>
      <div v-for="(card, idx) in s1Cards" :key="idx">
        <Card class="rounded-t-3xl rounded-b-3xl bg-black">
          <CardHeader class="rounded-t-3xl bg-gray-100">
            <CardTitle>{{ card.title }}</CardTitle>
            <CardDescription>{{ card.description }}</CardDescription>
          </CardHeader>
          <CardContent class="bg-gray-100"> {{ card.content }} </CardContent>
          <CardFooter class="rounded-b-3xl bg-gray-100"> {{ card.footer }} </CardFooter>
        </Card>
      </div>
    </GridLayout>
  </div>
</template>
