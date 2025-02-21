import { Response } from "@/types/response";

class Request {
    private headers = { 'Content-Type': 'application/json' };

    constructor(private baseUrl: string) {}

    private async request<T>(url: string, method: string, data?: any): Promise<Response<T>> {
        const options: RequestInit = {
            method,
            headers: this.headers,
            body: data ? JSON.stringify(data) : undefined,
        };
        const response = await fetch(`${this.baseUrl}${url}`, options);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
    }

    public get<T>(url: string, params?: Record<string, string>): Promise<Response<T>> {
        const queryString = params ? `?${new URLSearchParams(params).toString()}` : '';
        return this.request<T>(`${url}${queryString}`, 'GET');
    }

    public post<T>(url: string, data: any): Promise<Response<T>> {
        return this.request<T>(url, 'POST', data);
    }

    public put<T>(url: string, data: any): Promise<Response<T>> {
        return this.request<T>(url, 'PUT', data);
    }

    public delete(url: string): Promise<Response<any>> {
        return this.request(url, 'DELETE');
    }
}

export default new Request('http://localhost:9001');
