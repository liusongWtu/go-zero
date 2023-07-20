import { ref, nextTick, type Ref } from 'vue';

import { FormInstance } from 'element-plus';

import { type FormDataType, searchOptionsToFormData, formDataToSearchOptions } from './formData';

import { usePageOptionsStoreHook, getDefaultSearchOptions } from '@/store/modules/schemaMgrPageOpts';

import debounce from 'lodash-es/debounce';

export const useForm = (refForm: Ref<FormInstance>, emit: (evt: 'search') => void) => {
    const pageOptsStore = usePageOptionsStoreHook();
    const searchOpts = pageOptsStore.searchOpts;

    const formData = ref<FormDataType>(searchOptionsToFormData(searchOpts));

    const onReset = () => {
        formData.value = searchOptionsToFormData(getDefaultSearchOptions());
        pageOptsStore.updateSearchOptions(formDataToSearchOptions(formData.value));
        pageOptsStore.recordState();
        emit('search');
    };

    const onSearch = () => {
        refForm.value?.validate(isValid => {
            if (isValid) {
                pageOptsStore.updateSearchOptions(formDataToSearchOptions(formData.value));
                pageOptsStore.recordState();
                emit('search');
            }
        });
    };

    const debounceSearch = debounce(onSearch, 600);

    const doDebounceSearch = _val => nextTick(() => debounceSearch());

    return {
        formData,
        onReset: debounce(onReset, 600),
        onSearch,
        doDebounceSearch
    };
};
