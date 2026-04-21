<script lang="ts">
  import './app.css';
  import { page } from '$app/stores'; // Untuk cek URL aktif
  import { ui } from '$lib/ui.svelte'; // Jika kamu sudah buat state global kemarin
  
  // Import icons
  import home from '$lib/assets/home_1946488.png';
  import kelas from '$lib/assets/community_5528475.png';
  import pengajar from '$lib/assets/teacher_4833100.png';
  import kalender from '$lib/assets/calendar_1835377.png';
  import pengaturan from '$lib/assets/settings_2099058.png';
  import bantuan from '$lib/assets/question_471664.png';
  import profil from '$lib/assets/user_1946429.png';
  import plus from '$lib/assets/more_446136.png';

  let isSidebarExpanded = $state(false);
  let { children } = $props();

  // Mapping judul halaman untuk Breadcrumb
  const pageTitles: Record<string, string> = {
    '/': 'Beranda',
    '/kelas': 'Kelas',
    '/pengajar': 'Pengajar',
    '/kalender': 'Kalender',
    '/pengaturan': 'Pengaturan',
    '/bantuan': 'Bantuan',
    '/profil': 'Profil'
  };

  // Logika breadcrumb dinamis
  let currentTitle = $derived(pageTitles[$page.url.pathname] || 'Detail');

  const menuItems = [
    { img: home, label: 'Beranda', path: '/' },
    { img: kelas, label: 'Kelas', path: ' /kelas', hasSub: true },
    { img: pengajar, label: 'Pengajar', path: '/pengajar', hasSub: true },
    { img: kalender, label: 'Kalender', path: '/kalender' },  
    { img: pengaturan, label: 'Pengaturan', path: '/pengaturan' },
    { img: bantuan, label: 'Bantuan', path: '/bantuan' },
    { img: profil, label: 'Profil', path: '/profil' }
  ];
</script>

<div class="navbar bg-white border-b border-gray-100 px-4 h-16 fixed top-0 w-full z-[100]">
  <div class="flex flex-1 items-center  gap-2">
    <button onclick={() => isSidebarExpanded = !isSidebarExpanded} class="btn btn-square btn-ghost">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-6 h-6 stroke-gray-700"> 
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path> 
      </svg>
    </button>
    
    <div class="flex items-center gap-2 pl-2">
        <span class="text-xl font-bold tracking-tight text-black">PATH</span>
        <span class="text-gray-300 font-light text-xl">|</span>
        <span class="text-sm font-bold text-gray-500 mt-1 capitalize">{currentTitle}</span>
    </div>
  </div>

   <button class="btn btn-ghost btn-circle btn-sm border border-gray-300 mr-4">
        <img src="{plus}" alt="tambah">
    </button>

  <div class="flex-none gap-2">
    <div class="avatar">
      <div class="w-10 rounded-full ring-1 ring-gray-100">
        <img src="https://img.daisyui.com/images/profile/demo/yellingcat@192.webp" alt="profile" />
      </div>
    </div>
  </div>
</div>

<div class="flex h-[calc(100vh-64px)] pt-16 overflow-hidden bg-white">
  
  <aside 
  class="bg-white border-r border-gray-100 flex flex-col transition-all duration-300 ease-in-out z-[70]
    /* Mobile: Melayang (Overlay) */
    fixed inset-y-0 left-0 transform 
    {isSidebarExpanded ? 'translate-x-0 w-64' : '-translate-x-full w-20 md:translate-x-0'} 
    /* Desktop: Sejajar */
    md:relative md:translate-x-0"
>
   <div class="flex-1 overflow-y-auto overflow-x-hidden p-3 no-scrollbar mt-16 md:mt-0">

      <ul class="menu w-full p-0 gap-2">
        
        {#each menuItems as item}
          {@const isActive = $page.url.pathname === item.path || ($page.url.pathname.startsWith(item.path) && item.path !== '/')}
          <li>
            <a 
              href={item.path}
              class="flex items-center p-3 rounded-xl transition-all duration-200 group
              {isActive ? 'bg-[#EBE5FF] text-[#6338F1]' : 'text-gray-600 hover:bg-gray-50'}"
            >
              <img 
                src={item.img} 
                alt={item.label} 
                class="w-6 h-6 min-w-[24px] {isActive ? '' : 'grayscale opacity-70'}" 
              />
              
              {#if isSidebarExpanded}
                <span class="ml-4 font-semibold whitespace-nowrap flex-1 text-left">
                    {item.label}
                </span>
                {#if item.hasSub}
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-2 opacity-40" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                {/if}
              {/if}
            </a>
          </li>
        {/each}

      </ul>
    </div>
  </aside>

  {#if isSidebarExpanded}
  <div 
    onclick={() => isSidebarExpanded = false}
    class="fixed inset-0 bg-black/20 backdrop-blur-sm z-[60] md:hidden"
  ></div>
{/if}

  <main class="flex-1 p-2 md:p-4 bg-white relative w-full">
  <div class="bg-[#F8FAFC] w-full h-full rounded-[2rem] md:rounded-[2.5rem] border border-gray-100 overflow-auto">
    {@render children()}
  </div>
</main>
</div>

<style>
  .no-scrollbar::-webkit-scrollbar {
    display: none;
  }
</style>