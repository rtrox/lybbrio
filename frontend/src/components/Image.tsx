import { useState, useEffect } from "react";
import { useAuth } from "../context/AuthProvider";

async function convertBlobToBase64(blob: Blob) {
  const reader = new FileReader();
  reader.readAsDataURL(blob);
  return new Promise<string>((resolve, reject) => {
    reader.onload = () => {
      if (reader.result) {
        resolve(reader.result.toString());
      } else {
        reject("Failed to convert blob to base64");
      }
    };
  });
}

interface ImgProps extends React.ImgHTMLAttributes<HTMLImageElement> {
  url: string;
  alt: string;
}

export function Img({ url, alt, ...props }: ImgProps) {
  const { fetcher } = useAuth();
  const [data, setData] = useState<string>("");
  useEffect(() => {
    interface BlobAndType {
      blob: Blob;
      contentType: string | null;
    }

    fetcher(url, {
      method: "GET",
    })
      .then(async (res: Response) => {
        const b = await res.blob();
        return {
          blob: b,
          contentType: res.headers.get("content-type"),
        };
      })
      .then(async (res: BlobAndType) => {
        const b64data = await convertBlobToBase64(res.blob);
        setData(b64data);
      })
      .catch((error: any) => {
        console.error("Image fetch failed: ", error);
      });
  }, [fetcher]);
  return <img src={data} alt={alt} {...props} />;
}
