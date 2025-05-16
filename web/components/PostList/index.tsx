"use client";

import { useEffect, useState } from "react";
import { getPostDetailListAPI } from "@/api/post";
import { PostCard } from "@/components/PostCard";
import { Pagination } from "@heroui/pagination";
import type { PostVO } from "@/types/post";
import type { PageVO, Page } from "@/types/page";

export function PostList({ initialData }: { initialData: PageVO<PostVO> }) {
    const [pageNo, setPageNo] = useState(initialData.page_no);
    const [totalPage, setTotalPage] = useState(initialData.total_page);
    const [posts, setPosts] = useState<PostVO[]>(initialData.list);

    const fetchPosts = async (page_no: number) => {
        const req: Page = {
            page_no,
            page_size: 10,
        };
        const res = await getPostDetailListAPI(req);
        setPosts(res?.data?.list || []);
        setTotalPage(res?.data?.total_page || 1);
    };

    const updatePageNo = (page: number) => {
        setPageNo(page);
        fetchPosts(page);
    };

    useEffect(() => {
        // 避免重复请求首屏数据
        if (pageNo !== initialData.page_no) {
            fetchPosts(pageNo);
        }
    }, [pageNo]);

    return (
        <>
            {posts.map((item) => (
                <PostCard className="w-full" key={item.id} post={item} />
            ))}
            {totalPage > 1 && (
                <div className="my-10 flex justify-end">
                    <Pagination
                        initialPage={pageNo}
                        total={totalPage}
                        onChange={updatePageNo}
                    />
                </div>
            )}
        </>
    );
}
