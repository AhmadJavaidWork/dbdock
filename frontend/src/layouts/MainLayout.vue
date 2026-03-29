<template>
  <div class="flex h-screen">
    <!-- Sidebar -->
    <aside class="w-60 bg-gray-900 text-white p-4">
      <h2 class="font-bold mb-4">DBDock</h2>
      <!-- later: connections list -->
    </aside>

    <!-- Main content -->
    <main class="flex-1 p-4 dark:bg-background-dark text-text-light dark:text-text-dark">
      <RouterView />
    </main>
  </div>
</template>
