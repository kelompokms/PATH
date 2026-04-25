<script lang="ts">
  import { format, startOfWeek, addDays, isSameDay, addWeeks, subWeeks } from 'date-fns';
  import { id } from 'date-fns/locale'; // Untuk nama hari/bulan bahasa Indonesia

  // State untuk tanggal referensi (minggu yang sedang dilihat)
  let viewDate = $state(new Date());

  // Derived state: Menghitung 7 hari dalam minggu tersebut secara otomatis
  let weekDays = $derived.by(() => {
    // Mencari hari Senin di minggu tersebut
    let start = startOfWeek(viewDate, { weekStartsOn: 1 }); 
    return Array.from({ length: 7 }, (_, i) => addDays(start, i));
  });

  // Fungsi Navigasi
  const nextWeek = () => viewDate = addWeeks(viewDate, 1);
  const prevWeek = () => viewDate = subWeeks(viewDate, 1);

  // Data Dummy Jadwal (Nantinya bisa kamu ambil dari database/API)
  const jadwal = [
    { tanggal: new Date(2026, 3, 20), judul: 'Mapel 5: Skripsi', jam: '01.00' }, // Contoh 20 April 2026
  ];

  // Helper untuk mencari jadwal di tanggal tertentu
  const getJadwalByDate = (date: Date) => {
    return jadwal.filter(item => isSameDay(item.tanggal, date));
  };
</script>

<div class="flex flex-col h-full w-full p-4 md:p-6 bg-white">
  
  <div class="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
    <div class="dropdown">
      <div tabindex="0" role="button" class="btn bg-white border-gray-200 rounded-full px-6 normal-case shadow-sm hover:bg-gray-50">
        Semua Kelas
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
        </svg>
      </div>
      <ul tabindex="0" class="dropdown-content z-[1] menu p-2 shadow-lg bg-base-100 rounded-box w-52 mt-2">
        <li><button>Informatika A</button></li>
        <li><button>Informatika B</button></li>
      </ul>
    </div>

    <div class="flex items-center gap-4">
      <button onclick={prevWeek} class="btn btn-circle btn-sm bg-[#EBE5FF] border-none text-[#6338F1] hover:bg-[#d8ccff]">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
        </svg>
      </button>
      
      <span class="font-semibold text-gray-700 min-w-[180px] text-center">
        {format(weekDays[0], 'd MMM', { locale: id })} - {format(weekDays[6], 'd MMM, yyyy', { locale: id })}
      </span>

      <button onclick={nextWeek} class="btn btn-circle btn-sm bg-[#EBE5FF] border-none text-[#6338F1] hover:bg-[#d8ccff]">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
      </button>
    </div>
  </div>

  <div class="hidden md:grid grid-cols-7 border-t border-gray-100 h-full">
    {#each weekDays as day}
      {@const items = getJadwalByDate(day)}
      {@const isToday = isSameDay(day, new Date())}

      <div class="border-r border-gray-50 last:border-none flex flex-col items-center py-4 min-h-[400px]">
        <span class="text-gray-400 text-sm mb-2">{format(day, 'EEEE', { locale: id })}</span>
        
        <div class="w-10 h-10 flex items-center justify-center rounded-full text-xl font-bold mb-6 
          {isToday ? 'bg-[#EBE5FF] text-[#6338F1]' : 'text-gray-800'}">
          {format(day, 'd')}
        </div>

        <div class="w-full px-2 space-y-3">
          {#each items as j}
            <div class="bg-white border border-gray-100 rounded-xl p-3 shadow-sm hover:shadow-md transition-shadow cursor-pointer">
              <p class="text-[11px] font-bold text-gray-800 leading-tight">{j.judul}</p>
              <p class="text-[10px] text-gray-400 mt-1">{j.jam}</p>
            </div>
          {/each}
        </div>
      </div>
    {/each}
  </div>

  <div class="md:hidden flex flex-col gap-6">
    {#each weekDays as day}
      {@const items = getJadwalByDate(day)}
      {@const isToday = isSameDay(day, new Date())}
      
      <div class="flex items-start gap-4 border-b border-gray-100 pb-4">
        <div class="flex flex-col items-center min-w-[50px]">
          <span class="text-xs text-gray-400">{format(day, 'EEE', { locale: id })}</span>
          <div class="w-10 h-10 flex items-center justify-center rounded-full font-bold mt-1
            {isToday ? 'bg-[#EBE5FF] text-[#6338F1]' : 'text-gray-800'}">
            {format(day, 'd')}
          </div>
        </div>
        
        <div class="flex-1 space-y-2">
          {#each items as j}
            <div class="bg-white border border-gray-100 rounded-xl p-3 shadow-sm flex justify-between items-center">
              <span class="text-sm font-semibold text-gray-700">{j.judul}</span>
              <span class="text-xs text-gray-400">{j.jam}</span>
            </div>
          {:else}
            <div class="h-10 flex items-center">
               <div class="w-full h-[1px] bg-gray-100"></div>
            </div>
          {/each}
        </div>
      </div>
    {/each}
  </div>

</div>