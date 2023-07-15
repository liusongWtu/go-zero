import { store } from '@/store';
import { createPageOptsStore } from '@/utils/storeHelper';

import { CommonStatus } from '@/defines';
import { type ListReq } from '@/api/config';
import { getStringDate, getTimeFromStringDateTime } from '@/utils/timeUtils';

const STORE_ID = '{{.upperStartCamelObject}}-page-options';
const LOCAL_NAME = '{{.upperStartCamelObject}}-pageOpts';

export interface SearchOptions {
    startTime?: number;
    endTime?: number;
    ids?: number[];
    name?: string;
    code?: string;
    status?: CommonStatus[];
}

export function getDefaultSearchOptions(): SearchOptions {
    return {
        startTime: Math.floor(getTimeFromStringDateTime(`${getStringDate(6)} 00:00:00`) / 1000),
        endTime: Math.floor(getTimeFromStringDateTime(`${getStringDate(0)} 23:59:59`) / 1000)
    };
}

export function searchOptions2ListReq(opts: SearchOptions): ListReq {
    const result: ListReq = {};

    opts.startTime && (result.start_time = opts.startTime);
    opts.endTime && (result.end_time = opts.endTime);
    opts.ids && (result.ids = opts.ids);
    opts.name && (result.name = opts.name);
    opts.code && (result.code = opts.code);
    opts.status && (result.status = opts.status);

    return result;
}

export const { usePageOptionsStore } = createPageOptsStore<SearchOptions>(STORE_ID, LOCAL_NAME, {});

export function usePageOptionsStoreHook() {
    return usePageOptionsStore(store);
}
