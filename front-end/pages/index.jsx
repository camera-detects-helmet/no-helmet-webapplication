export default function Home() {
  return (
    <>
      <p className="text-gray-700 text-3xl mb-16 font-bold">Dashboard</p>

      <div className="grid lg:grid-cols-3 gap-5 mb-16">
        <div className="rounded bg-white h-40 shadow-sm"></div>
        <div className="rounded bg-white h-40 shadow-sm"></div>
        <div className="rounded bg-white h-40 shadow-sm"></div>

        <div className="relative">
        <div className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
          <svg aria-hidden="true" className="w-5 h-5 text-gray-500 dark:text-orange-500" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
            <path fillRule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clipRule="evenodd" /></svg>
        </div>
        <input datepicker datepicker-buttons type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5  dark:bg-orange-100 dark:border-orange-600 dark:placeholder-orange-500  dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date" />
      </div>

      </div>
      <div className="grid col-1 bg-white mb-20 shadow-sm overflow-y-auto h-screen">
    
        <div className=" bg-white-50">
          {/* Header */}
        
          <section className="min-h-screen body-font text-gray-600 ">
            <div className="container mx-auto px-5 py-10">
              <div className="-m-4 flex flex-wrap">
              

                {
                  Array.from({ length: 200 }, (_, i) => i + 1).map((item) => (
                    <div className="w-full p-4 md:w-1/2 lg:w-1/4">
                      <a className="relative block h-48 overflow-hidden rounded">
                        <img alt="ecommerce" className="block h-full w-full object-cover object-center cursor-pointer" src="https://dummyimage.com/421x261" />
                      </a>
                      <div className="mt-4">
                        <h3 className="title-font mb-1 text-xs tracking-widest text-gray-500">PROJECT</h3>
                        <h2 className="title-font text-lg font-medium text-gray-900">{ item}</h2>
                        <p className="mt-1">01/09/2022</p>
                      </div>
                    </div>
                  ))
                }
               
             
              </div>
            </div></section>
          {/* Footer */}
          <footer>
            <p className="text-center  py-4 bg-white">Create By © YOTHIN INTHAVYXUP</p>
          </footer>
        </div>
      </div>
    </>
  );
}
