<script lang="ts">
import bantuan from '$lib/assets/question_471664.png';

  const helpItems = [
    {
      title: 'Panduan AI',
      content: 'Di sini kamu bisa menemukan panduan lengkap tentang cara berinteraksi dengan asisten AI di platform PATH, mulai dari cara bertanya hingga mengoptimalkan hasil jawaban.'
    },
    {
      title: 'Etika Penggunaan AI',
      content: 'Gunakan AI sebagai alat bantu belajar, bukan untuk menyontek. Pastikan untuk selalu memverifikasi informasi kritis dan gunakan bahasa yang sopan.'
    },
    {
      title: 'FAQ',
      content: 'Kumpulan pertanyaan yang sering diajukan mengenai teknis aplikasi, manajemen kelas, dan pengaturan akun.'
    },
    {
      title: 'Laporkan Kesalahan AI',
      content: 'Jika kamu menemukan jawaban AI yang tidak akurat atau tidak pantas, kamu bisa melaporkannya melalui tombol lapor di dalam chat atau melalui formulir ini.'
    }
  ];
</script>

<div class="w-full min-h-screen bg-[#F8FAFC] p-4 md:p-10 flex justify-center">
  
  <div class="w-full max-w bg-white rounded-[2rem] shadow-sm border border-gray-100 overflow-hidden self-start">
    
    <div class="bg-[#EBE5FF] p-5 md:px-12 md:py-8">
      <h2 class="text-xl md:text-2xl font-bold text-[#1E293B]">Bantuan</h2>
    </div>

    <div class="p-2 md:p-8">
      {#each helpItems as item}
        <div class="collapse collapse-arrow bg-white border-b border-gray-100 rounded-none last:border-none">
          <input type="radio" name="help-accordion" /> 
          
          <div class="collapse-title flex items-center gap-4 py-5 px-4 md:px-6">
            <div class="text-gray-400">
              <img src={bantuan} alt="Bantuan" class="w-6 md:w-8" />
            </div>
            <span class="text-sm md:text-lg font-semibold text-gray-700">
              {item.title}
            </span>
          </div>

          <div class="collapse-content px-14 md:px-16 text-gray-500"> 
            <p class="pb-6 text-sm md:text-base leading-relaxed">
              {item.content}
            </p>
          </div>
        </div>
      {/each}
    </div>

  </div>
</div>

<style>
  /* Resetting DaisyUI default rounding for a cleaner list look */
  .collapse {
    border-radius: 0 !important;
  }
  
  /* Menyesuaikan warna tanda panah agar tidak terlalu kontras */
  .collapse-arrow > .collapse-title:after {
    color: #94a3b8; 
  }
</style>