import React,{useState} from 'react'
import Datepicker from 'react-tailwindcss-datepicker'
const home = () => {

  const [date, setDate] = useState({
    startDate: new Date(),
    endDate: new Date().setMonth(new Date().getMonth() + 1),
  })

  const handleDateChange = (date) => {
    setDate(date)

    console.log(date)
  }

  return (
    <>
      <p className="text-gray-700 text-3xl mb-16 font-bold">Dashboard</p>

      <div className="grid lg:grid-cols-3 gap-5 mb-5">
        <div className="rounded bg-white h-40 shadow-sm"></div>
        <div className="rounded bg-white h-40 shadow-sm"></div>
        <div className="rounded bg-white h-40 shadow-sm"></div>

  

        
      </div>
      <div className="grid lg:grid-cols-3 gap-2 mb-2">
      
        <Datepicker
          value={date}
          onChange={handleDateChange}
          showShortcuts={true}
          shortcuts={{
            'Today': new Date(),
            'Yesterday': new Date().setDate(new Date().getDate() - 1),
            'Tomorrow': new Date().setDate(new Date().getDate() + 1),
            'Next 7 Days': new Date().setDate(new Date().getDate() + 7),
            'Next 30 Days': new Date().setDate(new Date().getDate() + 30),
            'This Month': new Date().setDate(new Date().getDate() + 30),
            'Next Month': new Date().setMonth(new Date().getMonth() + 1),
            'This Year': new Date().setFullYear(new Date().getFullYear()),
            'Next Year': new Date().setFullYear(new Date().getFullYear() + 1),
          }}
          showDateDisplay={true}
          showMonthArrow={true}
          showYearArrow={true}
          showMonthYearPickers={true}
          showWeekNumbers={true}
          showDayNameInCalendar={true}

        />



      </div>
      
      <div className="grid col-1 bg-white mb-20 shadow-sm overflow-y-auto h-screen">

        <div className=" bg-white-50">
          {/* Header */}

          <section className="min-h-screen body-font text-gray-600 ">
            <div className="container mx-auto px-5 py-10">
              <div className="-m-4 flex flex-wrap">


                {
                  Array.from({ length: 200 }, (_, i) => i + 1).map((item) => (
                    <div key={item} className="w-full p-4 md:w-1/2 lg:w-1/4">
                      <a className="relative block h-48 overflow-hidden rounded">
                        <img alt="ecommerce" className="block h-full w-full object-cover object-center cursor-pointer" src="https://dummyimage.com/421x261" />
                      </a>
                      <div className="mt-4">
                        <h3 className="title-font mb-1 text-xs tracking-widest text-gray-500">PROJECT</h3>
                        <h2 className="title-font text-lg font-medium text-gray-900">{item}</h2>
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
  )
}

export default home