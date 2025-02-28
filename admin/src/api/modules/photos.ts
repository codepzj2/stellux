import request from "@/utils/request";

const GetPhotos = () => {
    return request.get("file/image")
}