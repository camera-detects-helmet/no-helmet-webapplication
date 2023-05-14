import React, { useState, useEffect } from 'react'
import Datepicker from 'react-tailwindcss-datepicker'
import Link from 'next/link'
import Head from 'next/head'
import moment from 'moment/moment'

import { getAllImages } from '../services'

export default function home ({images}) {

  const [date, setDate] = useState({
    startDate: new Date(),
    endDate: new Date().setMonth(new Date().getMonth() + 1),
  })

  const handleDateChange = (date) => {
    setDate(date)

  }



  return (
    <>
      <Head>
        <title>NokNae</title>
      </Head>
      <p className="text-gray-700 text-3xl mb-16 font-bold">Dashboard</p>

      {/* <div className="grid lg:grid-cols-3 gap-5 mb-5">
        <div className="rounded bg-white h-40 shadow-sm">
          <div className="grid grid-rows-2 gap-2 p-10">
            <div>
              <p className='text-xl'>
                Today
              </p>
            </div>
            <div>
              <p className='text-center'>
                number of today
              </p>
            </div>
          </div>
        </div>
        <div className="rounded bg-white h-40 shadow-sm">
        <div className="grid grid-rows-2 gap-2 p-10">
            <div>
              <p className='text-xl'>
                This month
              </p>
            </div>
            <div>
              <p className='text-center'>
                number of this month
              </p>
            </div>
          </div>
        </div>
        <div className="rounded bg-white h-40 shadow-sm">
        <div className="grid grid-rows-2 gap-2 p-10">
            <div>
              <p className='text-xl'>
                Total
              </p>
            </div>
            <div>
              <p className='text-center'>
                number of Total
              </p>
            </div>
          </div>
        </div> 




       </div> */}
      {/* <div className="grid lg:grid-cols-3 gap-2 mb-2">

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



      </div> */}

      <div className="grid col-1 bg-gray-50 mb-20 shadow-sm overflow-y-auto h-screen">
        <div>
          {/* Product List */}
          <section className="py-10 bg-gray-50">
            <div className="mx-auto grid max-w-24xl grid-cols-1 gap-6 p-6 sm:grid-cols-1 md:grid-cols-2  lg:grid-cols-3 xl:grid-cols-4">


              {images.map((data, index) =>
                // <div className="flex h-screen items-center justify-center bg-indigo-50 px-4">

                // </div>
                <div className="overflow-hidden rounded-xl bg-white shadow-md duration-200 hover:scale-105 hover:shadow-xl" key={data.id}>
                  <img src={data.path_default_img} alt="plant" className="h-auto w-full" />
                  <div className="p-5">
                    <p className="text-medium mb-5 text-gray-700">{data.imgName}</p>
                    <p className="text-medium mb-5 text-gray-700">Location : {data.location}</p>
                    <p className="text-medium mb-5 text-gray-700">Date : {moment(data.createAt).format('MMMM Do YYYY, h:mm:ss a')}</p>
                    <Link href={`/content_detail/${data.id}`}>
                      <button className="w-full rounded-md bg-indigo-600  py-2 text-indigo-100 hover:bg-indigo-500 hover:shadow-md duration-75">See More</button>

                    </Link>
                  </div>
                </div>
              )}


            </div>
          </section>

        </div>



      </div>
    </>
   
  )
}



export async function getStaticProps() {
  const response = (await getAllImages()) || [];


  return {
    props: {
      images: response.data.data.data.reverse(),
    },

  }
}

