<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRepositoryStore } from '~/stores/Repositories'
import type { Repository } from '~/stores/Repositories'

const repoStore = useRepositoryStore()

onMounted(() => {
  repoStore.init()
})

const searchQuery = ref('')
const drawerOpen = ref(false)
const editingRepo = ref<Repository | null>(null)
const saving = ref(false)

const formData = ref({
  name: '',
  url: '',
  description: '',
  enabled: true,
  priority: 0,
})

const filteredRepositories = computed(() => {
  if (!searchQuery.value) return repoStore.sortedRepositories
  const query = searchQuery.value.toLowerCase()
  return repoStore.sortedRepositories.filter(
    (repo) =>
      repo.name.toLowerCase().includes(query) ||
      repo.url.toLowerCase().includes(query) ||
      repo.description.toLowerCase().includes(query),
  )
})

const openCreateDrawer = () => {
  editingRepo.value = null
  formData.value = { name: '', url: '', description: '', enabled: true, priority: 0 }
  drawerOpen.value = true
}

const openEditDrawer = (repo: Repository) => {
  editingRepo.value = repo
  formData.value = {
    name: repo.name,
    url: repo.url,
    description: repo.description,
    enabled: repo.enabled,
    priority: repo.priority,
  }
  drawerOpen.value = true
}

const closeDrawer = () => {
  drawerOpen.value = false
  editingRepo.value = null
}

const saveRepository = async () => {
  saving.value = true
  try {
    if (editingRepo.value) {
      await repoStore.updateRepository(
        editingRepo.value.id,
        formData.value.name,
        formData.value.url,
        formData.value.description,
        formData.value.enabled,
        formData.value.priority,
      )
    } else {
      await repoStore.addRepository(
        formData.value.name,
        formData.value.url,
        formData.value.description,
        formData.value.enabled,
        formData.value.priority,
      )
    }
    closeDrawer()
  } finally {
    saving.value = false
  }
}

const deleteRepository = async (id: string) => {
  await repoStore.removeRepository(id)
}

const toggleRepository = async (id: string) => {
  await repoStore.toggleRepository(id)
}

const pingRepository = (id: string) => {
  repoStore.pingRepository(id).catch(() => {})
}

const pingAllRepositories = () => {
  repoStore.pingAll().catch(() => {})
}

const getStatusColor = (id: string) => {
  const s = repoStore.getStatus(id)
  if (!s) return 'bg-muted-foreground'
  if (s.online) return 'bg-primary'
  return 'bg-destructive'
}

const getStatusText = (id: string) => {
  const s = repoStore.getStatus(id)
  if (!s) return 'Unknown'
  return s.online ? 'Online' : 'Offline'
}
</script>

<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
      <div>
        <h1 class="text-2xl font-bold text-foreground">Repositories</h1>
        <p class="text-sm text-muted-foreground">Manage modpack sources and repositories</p>
      </div>
      <div class="flex gap-2">
        <button
          class="flex items-center gap-2 rounded bg-secondary px-4 py-2 text-sm font-medium text-secondary-foreground transition-colors hover:bg-secondary/80"
          @click="pingAllRepositories"
        >
          <Icon name="lucide:refresh-cw" class="h-4 w-4" />
          Ping All
        </button>
        <button
          class="flex items-center gap-2 rounded bg-primary px-4 py-2 text-sm font-medium text-primary-foreground transition-colors hover:bg-primary/90"
          @click="openCreateDrawer"
        >
          <Icon name="lucide:plus" class="h-4 w-4" />
          Add Repository
        </button>
      </div>
    </div>

    <!-- Search -->
    <div class="relative">
      <Icon
        name="lucide:search"
        class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
      />
      <input
        v-model="searchQuery"
        type="text"
        placeholder="Search repositories..."
        class="w-full rounded border border-border bg-input py-2.5 pl-10 pr-4 text-sm text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
      />
    </div>

    <!-- Repository List -->
    <div class="space-y-3">
      <div
        v-for="repo in filteredRepositories"
        :key="repo.id"
        class="group rounded-lg border border-border bg-card p-4 transition-colors hover:border-primary/30"
        :class="{ 'opacity-60': !repo.enabled }"
      >
        <div class="flex items-start justify-between gap-4">
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-3">
              <div class="flex h-10 w-10 items-center justify-center rounded-lg bg-secondary">
                <Icon name="lucide:database" class="h-5 w-5 text-primary" />
              </div>
              <div class="min-w-0 flex-1">
                <div class="flex items-center gap-2">
                  <h3 class="font-semibold text-foreground truncate">{{ repo.name }}</h3>
                  <span
                    class="flex items-center gap-1.5 rounded-full px-2 py-0.5 text-xs"
                    :class="
                      repo.enabled
                        ? 'bg-primary/10 text-primary'
                        : 'bg-muted text-muted-foreground'
                    "
                  >
                    <span class="h-1.5 w-1.5 rounded-full" :class="getStatusColor(repo.id)" />
                    {{ getStatusText(repo.id) }}
                  </span>
                </div>
                <p class="text-sm text-muted-foreground truncate">{{ repo.url }}</p>
              </div>
            </div>
            <p class="mt-2 text-sm text-muted-foreground line-clamp-2">
              {{ repo.description }}
            </p>
            <div
              v-if="repoStore.getStatus(repo.id)"
              class="mt-2 text-xs text-muted-foreground"
            >
              <span class="flex items-center gap-1">
                <Icon name="lucide:clock" class="h-3.5 w-3.5" />
                Last checked:
                {{
                  new Date(repoStore.getStatus(repo.id)!.checkedAt).toLocaleTimeString()
                }}
              </span>
            </div>
          </div>

          <!-- Actions -->
          <div class="flex items-center gap-1">
            <button
              class="flex h-8 w-8 items-center justify-center rounded text-muted-foreground transition-colors hover:bg-secondary hover:text-foreground"
              title="Ping repository"
              @click="pingRepository(repo.id)"
            >
              <Icon name="lucide:refresh-cw" class="h-4 w-4" />
            </button>
            <button
              class="flex h-8 w-8 items-center justify-center rounded text-muted-foreground transition-colors hover:bg-secondary hover:text-foreground"
              :title="repo.enabled ? 'Disable repository' : 'Enable repository'"
              @click="toggleRepository(repo.id)"
            >
              <Icon
                :name="repo.enabled ? 'lucide:toggle-right' : 'lucide:toggle-left'"
                class="h-4 w-4"
                :class="repo.enabled ? 'text-primary' : ''"
              />
            </button>
            <button
              class="flex h-8 w-8 items-center justify-center rounded text-muted-foreground transition-colors hover:bg-secondary hover:text-foreground"
              title="Edit repository"
              @click="openEditDrawer(repo)"
            >
              <Icon name="lucide:pencil" class="h-4 w-4" />
            </button>
            <button
              class="flex h-8 w-8 items-center justify-center rounded text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive"
              title="Delete repository"
              @click="deleteRepository(repo.id)"
            >
              <Icon name="lucide:trash-2" class="h-4 w-4" />
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div
        v-if="filteredRepositories.length === 0"
        class="flex flex-col items-center justify-center rounded-lg border border-dashed border-border py-12"
      >
        <Icon name="lucide:database" class="h-12 w-12 text-muted-foreground/50" />
        <h3 class="mt-4 text-lg font-medium text-foreground">No repositories found</h3>
        <p class="mt-1 text-sm text-muted-foreground">
          {{ searchQuery ? 'Try a different search term' : 'Add a repository to get started' }}
        </p>
        <button
          v-if="!searchQuery"
          class="mt-4 flex items-center gap-2 rounded bg-primary px-4 py-2 text-sm font-medium text-primary-foreground transition-colors hover:bg-primary/90"
          @click="openCreateDrawer"
        >
          <Icon name="lucide:plus" class="h-4 w-4" />
          Add Repository
        </button>
      </div>
    </div>

    <!-- Drawer Overlay -->
    <Teleport to="body">
      <Transition name="fade">
        <div
          v-if="drawerOpen"
          class="fixed inset-0 z-40 bg-black/50"
          @click="closeDrawer"
        />
      </Transition>
    </Teleport>

    <!-- Drawer -->
    <Teleport to="body">
      <Transition name="slide">
        <div
          v-if="drawerOpen"
          class="fixed bottom-0 right-0 top-0 z-50 w-full max-w-md border-l border-border bg-card shadow-xl"
        >
          <div class="flex h-full flex-col">
            <!-- Drawer Header -->
            <div class="flex items-center justify-between border-b border-border px-6 py-4">
              <h2 class="text-lg font-semibold text-foreground">
                {{ editingRepo ? 'Edit Repository' : 'Add Repository' }}
              </h2>
              <button
                class="flex h-8 w-8 items-center justify-center rounded text-muted-foreground transition-colors hover:bg-secondary hover:text-foreground"
                @click="closeDrawer"
              >
                <Icon name="lucide:x" class="h-5 w-5" />
              </button>
            </div>

            <!-- Drawer Content -->
            <div class="flex-1 overflow-y-auto p-6">
              <form class="space-y-5" @submit.prevent="saveRepository">
                <!-- Name -->
                <div class="space-y-2">
                  <label class="text-sm font-medium text-foreground">Repository Name</label>
                  <input
                    v-model="formData.name"
                    type="text"
                    placeholder="My Repository"
                    required
                    class="w-full rounded border border-border bg-input px-3 py-2.5 text-sm text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
                  />
                </div>

                <!-- URL -->
                <div class="space-y-2">
                  <label class="text-sm font-medium text-foreground">Repository URL</label>
                  <input
                    v-model="formData.url"
                    type="url"
                    placeholder="https://repo.example.com"
                    required
                    class="w-full rounded border border-border bg-input px-3 py-2.5 text-sm text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
                  />
                  <p class="text-xs text-muted-foreground">
                    Enter the base URL of the modpack repository
                  </p>
                </div>

                <!-- Description -->
                <div class="space-y-2">
                  <label class="text-sm font-medium text-foreground">Description</label>
                  <textarea
                    v-model="formData.description"
                    placeholder="A brief description of this repository..."
                    rows="3"
                    class="w-full resize-none rounded border border-border bg-input px-3 py-2.5 text-sm text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
                  />
                </div>

                <!-- Enabled Toggle -->
                <div class="flex items-center justify-between rounded-lg border border-border p-4">
                  <div>
                    <p class="text-sm font-medium text-foreground">Enable Repository</p>
                    <p class="text-xs text-muted-foreground">
                      Include this repository when fetching modpacks
                    </p>
                  </div>
                  <button
                    type="button"
                    class="relative h-6 w-11 rounded-full transition-colors"
                    :class="formData.enabled ? 'bg-primary' : 'bg-muted'"
                    @click="formData.enabled = !formData.enabled"
                  >
                    <span
                      class="absolute left-0.5 top-0.5 h-5 w-5 rounded-full bg-white transition-transform"
                      :class="{ 'translate-x-5': formData.enabled }"
                    />
                  </button>
                </div>
              </form>
            </div>

            <!-- Drawer Footer -->
            <div class="flex gap-3 border-t border-border px-6 py-4">
              <button
                type="button"
                class="flex-1 rounded border border-border bg-transparent px-4 py-2.5 text-sm font-medium text-foreground transition-colors hover:bg-secondary"
                @click="closeDrawer"
              >
                Cancel
              </button>
              <button
                type="button"
                :disabled="saving"
                class="flex-1 rounded bg-primary px-4 py-2.5 text-sm font-medium text-primary-foreground transition-colors hover:bg-primary/90 disabled:opacity-50"
                @click="saveRepository"
              >
                {{ saving ? 'Saving...' : editingRepo ? 'Save Changes' : 'Add Repository' }}
              </button>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}

.slide-enter-from,
.slide-leave-to {
  transform: translateX(100%);
}
</style>