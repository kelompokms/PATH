<script lang="ts">
import placehold from '$lib/assets/placehold.jpg'; // Menggunakan global state
  // State untuk melacak tab mana yang sedang aktif
  let activeTab = $state('forum');

  const tabs = [
    { id: 'forum', label: 'Forum' },
    { id: 'tugas', label: 'Tugas Kelas' },
    { id: 'anggota', label: 'Anggota' },
    { id: 'chat', label: 'Chat AI' }
  ];
</script>
  
 <div class="flex flex-col w-full min-h-screen bg-[#F8FAFC]">
  
  <div class="p-4 md:p-8">
    <div 
      class="w-full h-48 md:h-64 rounded-[2rem] p-8 flex flex-col justify-end relative overflow-hidden shadow-lg bg-cover bg-center"
      style="background-image: url({placehold});"
    >
      <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/20 to-transparent"></div>
      
      <div class="relative z-10">
        <h1 class="text-3xl md:text-4xl font-extrabold text-white">Informatika - A</h1>
        <p class="text-white/90 mt-2 font-medium tracking-wide">Tahun Ajaran 2023/2024</p>
      </div>
    </div>
  </div>

  <div class="px-4 md:px-10 mb-6">
    <div class="tabs tabs-boxed bg-transparent gap-2 p-0 overflow-x-auto no-scrollbar flex-nowrap">
      {#each tabs as tab}
        <button 
          onclick={() => activeTab = tab.id}
          class="btn rounded-full px-8 normal-case border-none transition-all
          {activeTab === tab.id 
            ? 'bg-[#EBE5FF] text-[#6338F1] shadow-sm' 
            : 'bg-white text-gray-400 hover:bg-gray-50'}">
          {tab.label}
        </button>
      {/each}
    </div>
  </div>

  <div class="px-4 md:px-10 pb-10 flex-1">
    
    {#if activeTab === 'forum'}
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="hidden md:flex flex-col gap-4">
          <div class="bg-white p-6 rounded-[1.5rem] border border-gray-100 shadow-sm">
            <h3 class="font-bold text-gray-800 mb-2 text-sm uppercase tracking-wider">Mendatang</h3>
            <p class="text-xs text-gray-400">Wah, tidak ada tugas yang perlu segera dikumpulkan!</p>
          </div>
        </div>

        <div class="md:col-span-3 flex flex-col gap-4">
          <div class="bg-white p-4 rounded-[1.5rem] border border-gray-100 shadow-sm flex items-center gap-4">
            <div class="w-10 h-10 rounded-full bg-[#EBE5FF]"></div>
            <input type="text" placeholder="Bagikan sesuatu dengan kelas..." class="input w-full bg-gray-50 rounded-xl border-none focus:ring-1 focus:ring-[#6338F1] text-sm" />
          </div>

          <div class="bg-white p-6 rounded-[1.5rem] border border-gray-100 shadow-sm">
             <div class="flex items-center gap-3 mb-4">
                <div class="w-10 h-10 rounded-full bg-slate-200"></div>
                <div>
                   <p class="text-sm font-bold text-gray-800">Admin</p>
                   <p class="text-[10px] text-gray-400">12 Apr</p>
                </div>
             </div>
             <p class="text-sm text-gray-600 leading-relaxed">
                Halo semuanya! Selamat datang di kelas Informatika A. Silakan cek bagian Tugas Kelas untuk materi terbaru.
             </p>
          </div>
        </div>
      </div>

    {:else if activeTab === 'tugas'}
      <div class="max-w-4xl mx-auto flex flex-col gap-4">
         {#each ['Materi 1', 'Tugas Praktikum'] as item}
            <div class="bg-white p-5 rounded-2xl border border-gray-100 shadow-sm flex items-center justify-between hover:bg-gray-50 cursor-pointer transition-all">
               <div class="flex items-center gap-4">
                  <div class="p-3 bg-[#EBE5FF] rounded-xl text-[#6338F1]">
                     <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" /></svg>
                  </div>
                  <span class="font-semibold text-gray-700">{item}</span>
               </div>
               <span class="text-xs text-gray-400 font-medium italic">Diposting 14 Apr</span>
            </div>
         {/each}
      </div>

    {:else if activeTab === 'anggota'}
       <div class="max-w-4xl mx-auto bg-white rounded-[2rem] border border-gray-100 overflow-hidden shadow-sm">
          <div class="p-6 border-b border-gray-50 font-bold text-[#6338F1] text-xl">Pengajar</div>
          <div class="p-6 flex items-center gap-4 border-b border-gray-50">
             <div class="w-10 h-10 rounded-full bg-[#EBE5FF]"></div>
             <span class="font-medium text-gray-700">Dosen Informatika</span>
          </div>
       </div>

    {:else if activeTab === 'chat'}
       <div class="h-[500px] bg-white rounded-[2rem] border border-gray-100 shadow-sm flex flex-col p-6 items-center justify-center text-center">
          <div class="w-16 h-16 bg-[#EBE5FF] rounded-2xl mb-4 flex items-center justify-center text-[#6338F1]">
             <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" /></svg>
          </div>
          <h2 class="text-xl font-bold text-gray-800">Tanya AI PATH</h2>
          <p class="text-gray-400 text-sm mt-2 max-w-xs">Tanyakan apa saja seputar materi di kelas ini, AI akan membantu menjawabnya.</p>
          <button class="btn bg-[#6338F1] border-none text-white rounded-full mt-6 px-10 normal-case">Mulai Chat</button>
       </div>
    {/if}

  </div>
</div>

<style>
  /* Menghilangkan scrollbar pada tab mobile */
  .no-scrollbar::-webkit-scrollbar {
    display: none;
  }
  .no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
  }
</style>