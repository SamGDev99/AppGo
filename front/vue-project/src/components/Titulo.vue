<template>
  <header class="flex items-center justify-center gap-3 px-6 py-4 mb-6 rounded-2xl bg-white shadow animate-fade-in-up">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="w-10 h-10 text-blue-600"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
      stroke-width="2"
    >
      <path stroke-linecap="round" stroke-linejoin="round" d="M3 3v18h18" />
      <path stroke-linecap="round" stroke-linejoin="round" d="M9 9l3 3-3 3" />
    </svg>
    <h1 class="text-4xl font-bold text-gray-800">Informe de Acciones</h1>
  </header>
</template>

<style scoped>
@keyframes fade-in-up {
  0% {
    opacity: 0;
    transform: translateY(10px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in-up {
  animation: fade-in-up 0.6s ease forwards;
}
</style>
