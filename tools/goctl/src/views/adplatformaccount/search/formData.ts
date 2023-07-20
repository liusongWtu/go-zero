import { type SearchOptions } from '@/store/modules/adPlatformAccountPageOpts';
import { unixTimeFormat, strDateRangeArrayToUnixTimeObject } from '@/utils/timeUtils';
import { CommonStatus } from '@/defines';

import { trimChar } from '@/utils/stringUtils';

import trim from 'ramda/es/trim';
import join from 'ramda/es/join';
import split from 'ramda/es/split';
import map from 'ramda/es/map';
import filter from 'ramda/es/filter';
import compose from 'ramda/es/compose';

export interface FormDataType {
    date?: [string, string];
    ids?: string;
    name?: string;
    code?: string;
    status?: Array<CommonStatus>;
}

export function searchOptionsToFormData(opts: SearchOptions): FormDataType {
    return {
        date: [
            opts?.startTime ? unixTimeFormat(opts.startTime) : '',
            opts?.endTime ? unixTimeFormat(opts.endTime) : ''
        ],
        ids: opts?.ids ? join(',', opts.ids) : '',
        name: opts.name ?? '',
        code: opts.code ?? '',
        status: opts.status ?? []
    };
}

export function formDataToSearchOptions(formData: FormDataType): SearchOptions {
    const result: SearchOptions = {};

    const toNumArray = compose(
        map((v: string) => Number(trim(v))),
        filter(v => v !== ''),
        split(','),
        trimChar(','),
        trim
    );

    const requestTime = strDateRangeArrayToUnixTimeObject(formData.date);

    requestTime.startTime && (result.startTime = requestTime.startTime);
    requestTime.endTime && (result.endTime = requestTime.endTime);

    formData.ids && (result.ids = toNumArray(formData.ids));
    formData.name && (result.name = formData.name);
    formData.code && (result.code = formData.code);
    formData.status && (result.status = formData.status);

    return result;
}
