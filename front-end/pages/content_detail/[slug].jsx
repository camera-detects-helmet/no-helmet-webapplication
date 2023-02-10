import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import axios from "axios";

const Detail = ({post}) => {

    const router = useRouter();
    const slug = router.query.slug;
    
  const [responseData, setResponseData] = useState([])

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    var data = '';

    var config = {
      method: 'get',
      url: "http://192.168.1.40:8081" + '/car/'+slug,
      headers: {},
      data: data
    };


    axios(config)
      .then(function (response) {
        setResponseData(response.data.data.data)
      })
      .catch(function (error) {
        console.log(error);
      });

  }
  
    
    return (
      <>
        <p className="text-gray-700 text-3xl mb-16 font-bold">Detail: { slug }</p>

        <div className="grid col-1 bg-gray-50 mb-20 shadow-sm overflow-y-auto h-screen">
          {/* <p>{ encodeURIComponent(post)}</p> */}
        </div>
      </>
    );
}

export default Detail;

// // Fetch data at build time
// export async function getStaticProps({ params }) {
//   const data = await getPostDetails(params.slug);
//   return {
//     props: {
//       post: data,
//     },
//   };
// }

// // Specify dynamic routes to pre-render pages based on data.
// // The HTML is generated at build time and will be reused on each request.
// export async function getStaticPaths() {
//   const posts = await getAllPosts();
//   return {
//     paths: posts.map(({ node: { slug } }) => ({ params: { slug } })),
//     fallback: true,
//   };
// }