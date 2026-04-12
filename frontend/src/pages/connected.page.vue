<script setup lang="ts">
import BaseBorderButton from "@/components/buttons/BaseBorderButton.vue";
import { useToast } from "@/composables/useToast";
import { useConnectionStore } from "@/stores/connection.store";
import { useLoaderStore } from "@/stores/loader";
import { storeToRefs } from "pinia";
import { onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";

const router = useRouter();

const connectionId = parseInt(useRoute().params["id"].toString());
const { activeConnections, selectedConnection } = storeToRefs(useConnectionStore());

async function disconnect() {
  if (!selectedConnection.value) return;
  useLoaderStore().show();
  try {
    const message = await useConnectionStore().disconnectFromDatabase(selectedConnection.value.id);
    useToast(message, "success");
  } catch (error) {}
  useLoaderStore().hide();
  router.replace({ name: "Welcome" });
}

onMounted(async function () {
  await useConnectionStore().getActiveConnections();
  if (activeConnections.value.length === 0 || !connectionId) {
    router.replace({ name: "Welcome" });
    return;
  }
  selectedConnection.value = null;
  for (let i = 0; i < activeConnections.value.length; i++) {
    if (connectionId === activeConnections.value[i].id) {
      selectedConnection.value = activeConnections.value[i];
      break;
    }
  }
  if (!selectedConnection.value) {
    router.replace({ name: "Welcome" });
  }
});
</script>

<template>
  <div>Connected to id:{{ useConnectionStore().selectedConnection?.id }}</div>
  <BaseBorderButton @click="disconnect">Disconnect</BaseBorderButton>
</template>
