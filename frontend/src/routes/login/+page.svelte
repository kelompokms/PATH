<script>
	import bg1 from '$lib/assets/Anime_Classroom_01_w.jpg.jpeg';
	import bg2 from '$lib/assets/5f7c5624c1339bf02d942e99a3d53882f9f60e5b_2_1035x582.jpeg';

	function login(event) {
		event.preventDefault();
		let formData = new FormData(event.target);

		const res = fetch('http://10.181.14.192:3000/login', {
			method: 'POST',
			credentials: 'include',
			body: formData
		}).then((res) => {
			console.log(res);
		});
	}
</script>

<div class="flex min-h-screen items-center justify-center bg-[#F8FAFC] bg-center p-4 md:p-8">
	<img src={bg1} alt="bg1" class="absolute inset-0 h-full w-full object-cover opacity-20" />
	<div
		class="card w-full max-w-6xl overflow-hidden rounded-[3rem] bg-white shadow-2xl transition-all duration-300"
	>
		<div class="flex h-full flex-col md:flex-row">
			<div class="relative h-64 overflow-hidden md:h-auto md:w-1/2">
				<img src={bg2} alt="bg2" class="h-full w-full object-cover" />

				<div
					class="absolute top-8 left-1/2 z-10 flex -translate-x-1/2 gap-2 rounded-full bg-white/70 p-2 shadow-md backdrop-blur-sm md:top-12 md:right-12 md:left-auto md:translate-x-0"
				>
					<a
						href="/daftar"
						class="btn btn-sm rounded-full border-none bg-[#E5E7EB]/50 px-6 font-semibold text-gray-500 normal-case hover:bg-[#E5E7EB]"
						>Daftar</a
					>
					<button
						class="btn btn-sm rounded-full border-none bg-[#E0F2FE] px-6 font-bold text-gray-800 normal-case"
						>Masuk</button
					>
				</div>
			</div>

			<div
				class="flex flex-col items-center p-10 text-center transition-all md:w-1/2 md:items-start md:p-16 md:text-left"
			>
				<div class="mb-10 flex w-full flex-col items-center gap-3 md:mb-14 md:items-start">
					<h2 class="text-3xl font-extrabold text-gray-950 md:text-5xl">Hi Student</h2>
					<p class="text-base font-medium tracking-wide text-gray-600 md:text-lg">
						welcome to PATH
					</p>
				</div>

				<form onsubmit={login} class="mx-auto flex w-full max-w-md flex-col gap-6 md:mx-0 md:gap-8">
					<div class="form-control w-full">
						<input
							name="email"
							type="email"
							placeholder="Email"
							class="input input-bordered h-12 w-full rounded-full border-gray-200 bg-white px-6 text-sm shadow-inner focus:border-[#6338F1] focus:ring-1 focus:ring-[#6338F1] md:h-14 md:text-base"
						/>
					</div>
					<div class="form-control relative w-full">
						<input
							name="password"
							type="password"
							placeholder="Password"
							class="input input-bordered h-12 w-full rounded-full border-gray-200 bg-white px-6 text-sm shadow-inner focus:border-[#6338F1] focus:ring-1 focus:ring-[#6338F1] md:h-14 md:text-base"
						/>
						<a
							href="/reset-password"
							class="absolute top-1/2 right-6 -translate-y-1/2 text-xs font-medium text-[#94A3B8] hover:text-[#6338F1] md:text-sm"
							>Lupa password ?</a
						>
					</div>

					<button
						class="btn mt-6 h-12 w-full rounded-full border-none bg-[#E0F2FE] text-base font-bold text-gray-800 normal-case shadow-sm transition-all hover:bg-[#BAE6FD] md:mt-10 md:h-14 md:text-lg"
						>Login</button
					>
				</form>

				<p class="mx-auto mt-10 w-full max-w-md text-center text-sm text-gray-600 md:mx-0 md:mt-12">
					Belum punya akun? <a href="/daftar" class="font-semibold text-[#6338F1] hover:underline"
						>Daftar</a
					>
				</p>
			</div>
		</div>
	</div>
</div>
