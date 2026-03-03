<script setup lang="ts">
import { ref, computed } from 'vue'

interface Instance {
  id: string
  name: string
  version: string
  mcVersion: string
  lastPlayed: string
  playTime: string
  size: string
  image: string
}

const instances = ref<Instance[]>([
  {
    id: '1',
    name: 'TechCraft Ultimate',
    version: '3.2.1',
    mcVersion: '1.20.4',
    lastPlayed: '2 hours ago',
    playTime: '124h 32m',
    size: '2.4 GB',
    image: 'https://images.unsplash.com/photo-1493711662062-fa541adb3fc8?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8dmlkZW8lMjBnYW1lfGVufDB8fDB8fHwwhttps://images.unsplash.com/photo-1493711662062-fa541adb3fc8?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8dmlkZW8lMjBnYW1lfGVufDB8fDB8fHww?w=200&h=200&fit=crop',
  },
  {
    id: '2',
    name: 'Medieval Kingdoms',
    version: '2.1.0',
    mcVersion: '1.20.4',
    lastPlayed: '1 day ago',
    playTime: '56h 18m',
    size: '1.8 GB',
    image: 'https://images.unsplash.com/photo-1518709268805-4e9042af9f23?w=200&h=200&fit=crop',
  },
  {
    id: '3',
    name: 'SkyFactory 5',
    version: '5.0.3',
    mcVersion: '1.20.1',
    lastPlayed: '3 days ago',
    playTime: '89h 45m',
    size: '3.1 GB',
    image: 'https://images.unsplash.com/photo-1534088568595-a066f410bcda?w=200&h=200&fit=crop',
  },
])

const totalStorage = computed(() => {
  const total = instances.value.reduce((acc, inst) => {
    const size = parseFloat(inst.size.replace(' GB', ''))
    return acc + size
  }, 0)
  return total.toFixed(1)
})

const handlePlay = (instance: Instance) => {
  console.log('Play instance:', instance.name)
}

const handleDelete = (id: string) => {
  instances.value = instances.value.filter((i) => i.id !== id)
}
</script>

<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <h1 class="font-display text-2xl font-bold text-foreground">
        Installed Instances
      </h1>
      <div class="flex items-center gap-2 text-sm text-muted-foreground">
        <Icon name="lucide:hard-drive" class="h-4 w-4" />
        <span>{{ totalStorage }} GB used</span>
      </div>
    </div>

    <!-- Instances Grid -->
    <div v-if="instances.length > 0" class="grid grid-cols-1 gap-4 lg:grid-cols-2">
      <div
        v-for="instance in instances"
        :key="instance.id"
        class="card flex gap-4 p-4"
      >
        <img
          :src="instance.image"
          :alt="instance.name"
          class="h-24 w-24 flex-shrink-0 rounded-lg object-cover"
        />
        <div class="flex flex-1 flex-col justify-between">
          <div>
            <div class="flex items-start justify-between">
              <h3 class="font-display text-lg font-semibold text-foreground">
                {{ instance.name }}
              </h3>
              <span class="rounded bg-secondary px-2 py-0.5 text-xs text-muted-foreground">
                {{ instance.mcVersion }}
              </span>
            </div>
            <p class="mt-1 text-sm text-muted-foreground">
              v{{ instance.version }}
            </p>
          </div>
          <div class="mt-3 flex items-center justify-between">
            <div class="flex items-center gap-4 text-xs text-muted-foreground">
              <span class="flex items-center gap-1">
                <Icon name="lucide:clock" class="h-3 w-3" />
                {{ instance.playTime }}
              </span>
              <span class="flex items-center gap-1">
                <Icon name="lucide:hard-drive" class="h-3 w-3" />
                {{ instance.size }}
              </span>
            </div>
            <div class="flex items-center gap-2">
              <button
                class="btn btn-ghost h-8 w-8 p-0 text-accent hover:bg-accent/10"
                aria-label="Delete instance"
                @click="handleDelete(instance.id)"
              >
                <Icon name="lucide:trash-2" class="h-4 w-4" />
              </button>
              <button
                class="btn btn-primary h-8 px-3 text-xs"
                @click="handlePlay(instance)"
              >
                <Icon name="lucide:play" class="h-3 w-3" />
                Play
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div
      v-else
      class="card flex flex-col items-center justify-center py-16 text-center"
    >
      <Icon name="lucide:folder-open" class="h-16 w-16 text-muted-foreground" />
      <h3 class="mt-4 text-lg font-medium text-foreground">
        No instances installed
      </h3>
      <p class="mt-2 text-muted-foreground">
        Browse modpacks and install one to get started
      </p>
    </div>
  </div>
</template>
