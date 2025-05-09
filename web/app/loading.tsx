
import { Spinner } from "@heroui/spinner";
export default function Loading() {
    return (
        <div className="flex justify-center items-center h-[70vh]">
            <Spinner classNames={{ label: "text-foreground mt-4" }} label="加载中..." variant="wave" />
        </div>
    )
}