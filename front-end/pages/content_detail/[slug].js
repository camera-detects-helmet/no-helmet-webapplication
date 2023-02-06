import { useRouter } from "next/router";

const Detail = () => {

    const router = useRouter();
    const { pid } = router.query;
    
    
    return (
      <>
        <p className="text-gray-700 text-3xl mb-16 font-bold">Detail:{pid}</p>

        <div className="grid col-1 bg-gray-50 mb-20 shadow-sm overflow-y-auto h-screen">
          <p>Test</p>
        </div>
      </>
    );
}

export default Detail;

