import { http } from '@/utils/http';
import { baseUrlApi } from './utils';
import { type CommonStatus } from '@/defines';
import { type AxiosRequestConfig } from 'axios';

export interface AdPlatformAccountItem {
        id?: number;  
    code?: string; // 平台账号code 
    name?: string; // 平台账号名称 
    platform_code?: string; // 广告平台code 
    platform_title?: string; // 广告平台title 
    data?: string; // 配置信息 
    status?: CommonStatus; // 状态，1=正常，2=冻结 
    create_time?: number; // 创建时间秒 
    update_time?: number; // 修改时间秒 
}

export type CreateReq = AdPlatformAccountItem;

export type UpdateReq = AdPlatformAccountItem;

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
    return http.request<Resp<AdPlatformAccountItem>>('post', baseUrlApi('/ad/platform/account/create'), config);
};

export const fetchUpdateConfig = (config: AxiosRequestConfig<UpdateReq>) => {
    return http.request<Resp<AdPlatformAccountItem>>('post', baseUrlApi('/ad/platform/account/update'), config);
};

export const fetchDeleteConfig = (config: AxiosRequestConfig<DeleteReq>) => {
    return http.request<Resp<null>>('post', baseUrlApi('/ad/platform/account/delete'), config);
};

export const fetchList = (config: AxiosRequestConfig<ListReq>) => {
    return http.request<Resp<RespList<AdPlatformAccountItem>>>('post', baseUrlApi('/ad/platform/account/list'), config);
};

export const fetchAdPlatformAccountItem = (config: AxiosRequestConfig<ItemReq>) => {
    return fetchList({ ...config, data: { ids: [config.data.id] } }).then(res => {
        if (res.data?.list && res.data.list.length > 0) {
            const result: Resp<AdPlatformAccountItem> = { code: 0, data: res.data.list[0], msg: '' };
            return Promise.resolve(result);
        } else {
            return Promise.reject('Record does not exist');
        }
    });
};
