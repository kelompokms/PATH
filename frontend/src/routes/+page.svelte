<script lang="ts">
  // Jika kamu punya file gambarnya, import di sini
  import uns from '$lib/assets/UNS-LOGO.png';
  
</script>

<div class="flex flex-col items-center justify-center h-full w-full p-6">
  
  <div class="bg-white p-8 md:p-12 rounded-[2.5rem] shadow-sm flex flex-col items-center max-w-sm w-full transition-all hover:shadow-md">
    
    <div class="w-48 h-48 md:w-64 md:h-64 mb-6 flex items-center justify-center">
      <img 
        src={uns} 
        alt="Ilustrasi Kelas" 
        class="w-full h-full object-contain mix-blend-multiply opacity-80"
      />
    </div>

    <p class="text-gray-500 text-sm md:text-base font-medium text-center">
      Gabung Kelas atau Buat Kelas
    </p>
  </div>

  <div class="flex flex-row gap-4 mt-8 w-full max-w-sm justify-center">
    
    <button class="btn border-none bg-[#DDEBFA] hover:bg-[#C5DFF8] text-[#1E3A8A] px-8 rounded-xl normal-case font-semibold flex-1 h-12 shadow-sm">
      Gabung Kelas
    </button>

    <button class="btn border-none bg-[#F1F5F9] hover:bg-[#E2E8F0] text-gray-600 px-8 rounded-xl normal-case font-semibold flex-1 h-12 shadow-sm">
      Buat Kelas
    </button>
    
  </div>

</div>

<style>
  /* Animasi halus saat halaman dimuat */
  div {
    animation: fadeIn 0.5s ease-out;
  }

  @keyframes fadeIn {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
  }
</style>