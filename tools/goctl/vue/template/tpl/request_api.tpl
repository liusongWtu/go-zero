import { http } from '@/utils/http';
import { baseUrlApi } from './utils';
import { type CommonStatus } from '@/defines';
import { type AxiosRequestConfig } from 'axios';

export interface {{.upperStartCamelObject}}Item {
    {{.fields}}
}

export type CreateReq = {{.upperStartCamelObject}}Item;

export type UpdateReq = {{.upperStartCamelObject}}Item;

export interface DeleteReq {
    ids?: Array<number>;
}

export interface ListReq extends BaseReqList {
    start_time?: number;
    end_time?: number;
    ids?: Array<number>;
    name?: string;
    code?: string;
    status?: Array<CommonStatus>;
}

export interface ItemReq {
    id?: number;
}

export const fetchCreateConfig = (config: AxiosRequestConfig<CreateReq>) => {
    return http.request<Resp<{{.upperStartCamelObject}}Item>>('post', baseUrlApi('/{{.requestPath}}/create'), config);
};

export const fetchUpdateConfig = (config: AxiosRequestConfig<UpdateReq>) => {
    return http.request<Resp<{{.upperStartCamelObject}}Item>>('post', baseUrlApi('/{{.requestPath}}/update'), config);
};

export const fetchDeleteConfig = (config: AxiosRequestConfig<DeleteReq>) => {
    return http.request<Resp<null>>('post', baseUrlApi('/{{.requestPath}}/delete'), config);
};

export const fetchList = (config: AxiosRequestConfig<ListReq>) => {
    return http.request<Resp<RespList<{{.upperStartCamelObject}}Item>>>('post', baseUrlApi('/{{.requestPath}}/list'), config);
};

export const fetch{{.upperStartCamelObject}}Item = (config: AxiosRequestConfig<ItemReq>) => {
    return fetchList({ ...config, data: { ids: [config.data.id] } }).then(res => {
        if (res.data?.list && res.data.list.length > 0) {
            const result: Resp<{{.upperStartCamelObject}}Item> = { code: 0, data: res.data.list[0], msg: '' };
            return Promise.resolve(result);
        } else {
            return Promise.reject('Record does not exist');
        }
    });
};
