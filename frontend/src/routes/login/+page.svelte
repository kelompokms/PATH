<script lang="ts">
  
</script>

<div class="min-h-screen flex items-center justify-center p-4 md:p-8 bg-[#F8FAFC]" style="background-image: url(); background-size: cover; background-position: center;">
  
  <div class="card bg-white shadow-2xl rounded-[3rem] overflow-hidden w-full max-w-6xl transition-all duration-300">
    <div class="flex flex-col md:flex-row h-full">
      
      <div class="md:w-1/2 h-64 md:h-auto relative overflow-hidden">
        <img src=""  alt="Classroom" class="w-full h-full object-cover md:rounded-[3rem] md:rounded-r-none rounded-t-[3rem] md:rounded-t-none" />
        
        <div class="absolute top-8 left-1/2 -translate-x-1/2 md:top-12 md:left-auto md:right-12 md:translate-x-0 flex gap-2 p-2 bg-white/70 backdrop-blur-sm rounded-full z-10 shadow-md">
            <a href="/daftar" class="btn btn-sm rounded-full normal-case text-gray-500 bg-[#E5E7EB]/50 hover:bg-[#E5E7EB] border-none font-semibold px-6">Daftar</a>
            <button class="btn btn-sm rounded-full normal-case text-gray-800 bg-[#E0F2FE] border-none font-bold px-6">Masuk</button>
        </div>
      </div>

      <div class="md:w-1/2 p-10 md:p-16 flex flex-col items-center md:items-start text-center md:text-left transition-all">
        <div class="w-full flex flex-col items-center md:items-start gap-3 mb-10 md:mb-14">
            <h2 class="text-3xl md:text-5xl font-extrabold text-gray-950">Hi Student</h2>
            <p class="text-base md:text-lg text-gray-600 font-medium tracking-wide">welcome to PATH</p>
        </div>
        
        <form class="w-full flex flex-col gap-6 md:gap-8 max-w-md mx-auto md:mx-0">
          <div class="form-control w-full">
            <input type="email" placeholder="Email" class="input input-bordered w-full rounded-full bg-white border-gray-200 focus:border-[#6338F1] focus:ring-1 focus:ring-[#6338F1] text-sm md:text-base h-12 md:h-14 px-6 shadow-inner" />
          </div>
          <div class="form-control w-full relative">
            <input type="password" placeholder="Password" class="input input-bordered w-full rounded-full bg-white border-gray-200 focus:border-[#6338F1] focus:ring-1 focus:ring-[#6338F1] text-sm md:text-base h-12 md:h-14 px-6 shadow-inner" />
            <a href="/reset-password" class="text-xs md:text-sm text-[#94A3B8] hover:text-[#6338F1] absolute right-6 top-1/2 -translate-y-1/2 font-medium">Lupa password ?</a>
          </div>
          
          <button class="btn bg-[#E0F2FE] hover:bg-[#BAE6FD] text-gray-800 border-none w-full rounded-full mt-6 md:mt-10 normal-case font-bold h-12 md:h-14 text-base md:text-lg shadow-sm transition-all">Login</button>
        </form>

        <p class="text-sm text-gray-600 mt-10 md:mt-12 text-center w-full max-w-md mx-auto md:mx-0">
          Belum punya akun? <a href="/daftar" class="text-[#6338F1] font-semibold hover:underline">Daftar</a>
        </p>
      </div>

    </div>
  </div>
</div>