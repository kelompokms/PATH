<script lang="ts">
  let profileData = $state({
    name: 'Siswa 1',
    email: 'siswa@mail.com',
    phone: '08123456789'
  });

  let isModalOpen = $state(false);
  let editingField = $state('');
  let tempValue = $state('');

  function openModal(field: string) {
    editingField = field;
    tempValue = profileData[field as keyof typeof profileData] || '';
    isModalOpen = true;
  }

  function saveChanges() {
    if (editingField in profileData) {
      profileData[editingField as keyof typeof profileData] = tempValue;
    }
    isModalOpen = false;
  }
</script>

<div class="flex flex-col items-center w-full min-h-screen p-4 md:p-8 bg-[#F8FAFC]">
  
  <div class="w-full max-w- bg-white rounded-[1.5rem] md:rounded-[2rem] shadow-sm border border-gray-100 overflow-hidden">
    
    <div class="bg-[#EBE5FF] p-4 md:px-10 md:py-6">
      <h2 class="text-xl md:text-2xl font-semibold text-gray-800">Profil</h2>
    </div>

    <div class="flex flex-col">
      <div class="flex flex-row justify-between items-center p-4 md:p-10 border-b border-gray-100">
        <div class="flex items-center gap-3 md:gap-6">
          <div class="w-12 h-12 md:w-20 md:h-20 rounded-full bg-white border-2 md:border-[3px] border-[#EBE5FF] flex items-center justify-center">
            <span class="text-[#6338F1]"><svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 md:h-12 md:w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" /></svg></span>
          </div>
          <span class="text-sm md:text-lg font-medium text-gray-700">Foto Profil</span>
        </div>
        <button class="btn btn-outline btn-sm md:btn-md border-gray-200 rounded-full px-4 md:px-8 normal-case text-gray-500">Edit</button>
      </div>

      {#each [
        { id: 'name', label: 'Nama Tampilan', value: profileData.name },
        { id: 'email', label: 'Email', value: profileData.email },
        { id: 'phone', label: 'Nomor Telepon', value: profileData.phone }
      ] as item}
        <div class="flex flex-row justify-between items-center p-4 md:p-10 border-b border-gray-100 last:border-none">
          <div class="flex-1 pr-2">
            <p class="text-[10px] md:text-sm font-semibold text-gray-400 uppercase tracking-wider">{item.label}</p>
            <p class="text-sm md:text-lg font-medium text-gray-800 mt-0.5 break-all">{item.value}</p>
          </div>
          <button 
            onclick={() => openModal(item.id)} 
            class="btn btn-outline btn-sm md:btn-md border-gray-200 rounded-full px-4 md:px-8 normal-case text-gray-500 shrink-0"
          >
            Edit
          </button>
        </div>
      {/each}
    </div>
  </div>
</div>

<input type="checkbox" id="profile_modal" class="modal-toggle" checked={isModalOpen} />
<div class="modal modal-bottom sm:modal-middle" role="dialog">
  <div class="modal-box bg-white rounded-t-[2rem] sm:rounded-[2rem] p-6">
    <h3 class="text-lg font-bold text-gray-800">Ubah Data</h3>
    <p class="text-sm text-gray-500 mb-4 italic">Masukkan {editingField} baru kamu di bawah ini.</p>
    
    <div class="form-control w-full">
      <input 
        type="text" 
        bind:value={tempValue} 
        placeholder="Ketik di sini..."
        class="input input-bordered w-full bg-gray-50 border-gray-200 focus:border-[#6338F1] focus:ring-1 focus:ring-[#6338F1] rounded-xl text-gray-800"
      />
    </div>

    <div class="modal-action flex flex-row gap-2">
      <button onclick={() => isModalOpen = false} class="btn btn-ghost flex-1 rounded-full normal-case">Batal</button>
      <button onclick={saveChanges} class="btn bg-[#6338F1] hover:bg-[#522ccb] text-white border-none flex-1 rounded-full normal-case">Simpan</button>
    </div>
  </div>
  <label class="modal-backdrop" for="profile_modal" onclick={() => isModalOpen = false}>Close</label>
</div>