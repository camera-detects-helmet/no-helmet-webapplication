import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { getImageById, getAllImages } from "../../services";

const Detail = ({ images }) => {

  const router = useRouter();
  const slug = router.query.slug;
  if (router.isFallback) {
    return <div>Loading...</div>;
  }



  return (
    <>
      <p className="text-gray-700 text-3xl mb-16 font-bold">Detail: {images.id}</p>

      <div className="grid col-1 bg-gray-50 mb-20 shadow-sm overflow-y-auto h-screen">

        <div className="w-full max-w-6xl rounded bg-white shadow-xl p-10 lg:p-20 mx-auto text-gray-800 relative md:text-left">
          <div className="md:flex items-center -mx-10">
            <div className="w-full md:w-1/2 px-10 mb-10 md:mb-0">
              <div className="relative m-2">
                <img src={images.path_default_img} className="w-full relative z-10" alt="" />
                <div className="border-4 border-yellow-200 absolute top-10 bottom-10 left-10 right-10 z-0" />
              </div>
              <div className="relative m-2">
                <img src={images.path_rider_img} className="w-full relative z-10" alt="" />
                <div className="border-4 border-yellow-200 absolute top-10 bottom-10 left-10 right-10 z-0" />
              </div>
            </div>

            <div className="w-full md:w-1/2 px-10">
              <div className="mb-10">
                <h1 className="font-bold  mb-5"> Image Name : {images.imgName}</h1>
                <p className="text-sm">Location : {images.location}</p>
                <p className="text-sm">Date : {images.createAt}</p>

              </div>
              <div>
                <div className="inline-block align-bottom">
                  {/* <button className="bg-yellow-300 opacity-75 hover:opacity-100 text-yellow-900 hover:text-gray-900 rounded-full px-10 py-2 font-semibold m-2 w-full"><i className="mdi mdi-cart -ml-2 mr-2" /> Download full-image </button> */}
                  {/* <button className="bg-yellow-300 opacity-75 hover:opacity-100 text-yellow-900 hover:text-gray-900 rounded-full px-10 py-2 font-semibold m-2 w-full"><i className="mdi mdi-cart -ml-2 mr-2" /> Download rider-imgage</button> */}

                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </>
  );
}

export default Detail;



export async function getStaticProps({ params }) {

  const response = await getImageById(params.slug);
  
  return {
    props: {
      images: response.data.data.data,
    },
  };
}

export async function getStaticPaths() {
  const response = await getAllImages();
  const res = response.data

  return {
    paths: res.data.data.map(image => ({ params: { slug: image.id } })),
    fallback: true,
  };
}

