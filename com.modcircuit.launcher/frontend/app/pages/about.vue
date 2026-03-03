<script setup lang="ts">
const version = '1.0.0'
const buildDate = 'January 31, 2026'


const links = [
  { label: 'Website', url: 'https://modcircuit.com', icon: 'lucide:globe' },
  { label: 'Discord', url: 'https://discord.gg/modcircuit', icon: 'lucide:message-circle' },
  { label: 'GitHub', url: 'https://github.com/modcircuit', icon: 'lucide:github' },
  { label: 'Twitter', url: 'https://twitter.com/modcircuit', icon: 'lucide:twitter' },
]

const handleCheckUpdate = () => {
  console.log('Checking for updates...')
}
</script>

<template>
  <div class="mx-auto max-w-2xl space-y-8">
    <!-- Logo and Title -->
    <div class="flex flex-col items-center text-center">
      <img
        src="/images/modcircuit-logo.png"
        alt="ModCircuit"
        class="h-32 w-32"
      />
      <h1 class="mt-4 font-display text-3xl font-bold text-foreground">
        ModCircuit Launcher
      </h1>
      <p class="mt-2 text-muted-foreground">
        Your ultimate Minecraft modpack launcher
      </p>
    </div>

    <!-- Version Info -->
    <div class="card p-6 text-center">
      <div class="grid grid-cols-2 gap-4">
        <div>
          <p class="text-sm text-muted-foreground">Version</p>
          <p class="mt-1 font-display text-lg font-semibold text-foreground">
            v{{ version }}
          </p>
        </div>
        <div>
          <p class="text-sm text-muted-foreground">Build Date</p>
          <p class="mt-1 font-display text-lg font-semibold text-foreground">
            {{ buildDate }}
          </p>
        </div>
      </div>
      <button class="btn btn-primary mt-6" @click="handleCheckUpdate">
        <Icon name="lucide:refresh-cw" class="h-4 w-4" />
        Check for Updates
      </button>
    </div>

    <!-- Links -->
    <div class="card p-6">
      <h2 class="mb-4 font-display text-lg font-semibold text-foreground">
        Connect With Us
      </h2>
      <div class="grid grid-cols-2 gap-3">
        <a
          v-for="link in links"
          :key="link.label"
          :href="link.url"
          target="_blank"
          rel="noopener noreferrer"
          class="btn btn-secondary justify-start"
        >
          <Icon :name="link.icon" class="h-4 w-4" />
          {{ link.label }}
        </a>
      </div>
    </div>

    <!-- Credits -->
    <div class="card p-6">
      <h2 class="mb-4 font-display text-lg font-semibold text-foreground">
        Credits
      </h2>
      <p class="text-sm text-muted-foreground">
        ModCircuit Launcher is built with love by the ModCircuit Team.
        Special thanks to the Minecraft modding community for their amazing work.
      </p>
      <p class="mt-4 text-sm text-muted-foreground">
        Minecraft is a trademark of Mojang AB. ModCircuit is not affiliated with
        or endorsed by Mojang AB.
      </p>
    </div>

    <!-- Footer -->
    <p class="text-center text-sm text-muted-foreground">
      Made with
      <Icon name="lucide:heart" class="inline h-4 w-4 text-accent" />
      by ModCircuit Team
    </p>
  </div>
</template>
