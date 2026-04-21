<script lang="ts">
  import notifikasi from '$lib/assets/notification_2645897.png';
  import brush from '$lib/assets/brush_951715.png';
  import { ui } from '$lib/ui.svelte'; // Menggunakan global state

  // State lokal untuk notifikasi
  let isNotifyActive = $state(false);

  // Fungsi toggle yang langsung mengubah tema global
  function handleThemeChange() {
    ui.theme = ui.theme === 'light' ? 'dark' : 'light';
  }
</script>

<div class="flex flex-col items-center w-full h-full p-4 md:p-10 bg-[#F8FAFC]">
  
  <div class="w-full max-w bg-white rounded-[2.5rem] shadow-sm border border-gray-100 overflow-hidden">
    
    <div class="bg-[#EBE5FF] px-8 py-6">
      <h2 class="text-2xl font-bold text-gray-800">Pengaturan</h2>
    </div>

    <div class="flex flex-col">
      
      <div class="flex justify-between items-center px-8 py-7 border-b border-gray-50 hover:bg-slate-50 transition-colors">
        <div class="flex items-center gap-5">
          <div class="p-3 ">
            <img src={notifikasi} alt="Notifikasi" class="w-8" />
          </div>
          <div>
            <p class="text-lg font-bold text-gray-800">Notifikasi</p>
          </div>
        </div>

        <div class="flex items-center gap-4 bg-[#EBE5FF] pl-6 pr-2 py-2 rounded-full min-w-[140px] justify-between transition-all">
          <span class="text-sm font-extrabold text-[#6338F1] uppercase tracking-wider">
            {isNotifyActive ? 'Aktif' : 'Mati'}
          </span>
          <input 
            type="checkbox" 
            bind:checked={isNotifyActive}
            class="toggle toggle-lg border-none bg-white [--tglbg:#6338F1] checked:bg-white rounded-full" 
          />
        </div>
      </div>

      <div class="flex justify-between items-center px-8 py-7 hover:bg-slate-50 transition-colors">
        <div class="flex items-center gap-5">
          <div class="p-3 rounded-2xl">
            <img src={brush} alt="Tema" class="w-8" />
          </div>
          <div>
            <p class="text-lg font-bold text-gray-800">Tema</p>
          </div>
        </div>

        <div class="flex items-center gap-4 bg-[#EBE5FF] pl-6 pr-2 py-2 rounded-full min-w-[140px] justify-between transition-all">
          <span class="text-sm font-extrabold text-[#6338F1] uppercase tracking-wider">
            {ui.theme === 'dark' ? 'Gelap' : 'Terang'}
          </span>
          <input 
            type="checkbox" 
            checked={ui.theme === 'dark'}
            onchange={handleThemeChange}
            class="toggle toggle-lg border-none bg-white [--tglbg:#6338F1] checked:bg-white rounded-full" 
          />
        </div>
      </div>

    </div>
  </div>

</div>

<style>
  :global(.toggle) {
    border-radius: 9999px;
    height: 2rem;
    width: 3.5rem;
  }
</style>