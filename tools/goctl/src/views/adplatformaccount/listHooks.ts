import {
    ref,
    reactive,
    computed,
    toRaw,
    createVNode,
    onMounted,
    onUnmounted,
    resolveDirective,
    withDirectives
} from 'vue';
import { useRouter, type RouteParamsRaw } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { type Sort } from 'element-plus';
import { type PaginationProps } from '@pureadmin/table';

import { JsonCodeMirror } from '@/components/JsonCodeMirror';
import { addDialog } from '@/components/ReDialog';
import { CustomDialogHF } from '@/common/CustomDialogHF';
import { AdPlatformAccountEditorForm, type AdPlatformAccountEditorFormInstance } from './editor/form';

import {
    createNormalFormatter,
    noWarpContent,
    createTimestampRender,
    createCommonStatusRender
} from '@/common/tableCellRenders';

import clone from 'ramda/es/clone';

import {
    fetchCreateAdPlatformAccount,
    fetchUpdateAdPlatformAccount,
    fetchDeleteAdPlatformAccount,
    fetchList,
    type AdPlatformAccountItem,
    type ListReq,
    type DeleteReq
} from '@/api/ad-platform-account';
import { usePageOptionsStoreHook, searchOptions2ListReq } from '@/store/modules/adPlatformAccountPageOpts';

import { useMultiTagsStoreHook } from '@/store/modules/multiTags';

import { CommonStatusActivate } from '@/defines';

import { message } from '@/utils/message';

import { useDeleteItemHook } from '@/common/ListHelpers/deleteItemHelper';

import { $t } from '@/plugins/i18n';

export function useList() {
    const { t } = useI18n();

    const renderHeader = data => noWarpContent(t(data.column.label));

    const statusRender = createCommonStatusRender({
        fieldName: 'status',
        fetchUpdate: fetchUpdateAdPlatformAccount,
        getUpdateParams: row => ({
            id: row.id,
            name: row.name,
            code: row.code,
            status: row.status
        }),
        getStatus: row => row.status,
        setStatus: (row, status) => (row.status = status),
        getName: row => row.name
    });

    const columns: TableColumnList = [
                {
            label: $t('adPlatformAccount.itemFields.id'),
            prop: 'id',
            minWidth: 100,
            renderHeader,
            formatter: createNormalFormatter('id')
        },
        {
            label: $t('adPlatformAccount.itemFields.code'),
            prop: 'code',
            minWidth: 100,
            renderHeader,
            formatter: createNormalFormatter('code')
        },
        {
            label: $t('adPlatformAccount.itemFields.name'),
            prop: 'name',
            minWidth: 100,
            renderHeader,
            formatter: createNormalFormatter('name')
        },
        {
            label: $t('adPlatformAccount.itemFields.platform_code'),
            prop: 'platform_code',
            minWidth: 100,
            renderHeader,
            formatter: createNormalFormatter('platform_code')
        },
        {
            label: $t('adPlatformAccount.itemFields.platform_title'),
            prop: 'platform_title',
            minWidth: 100,
            renderHeader,
            formatter: createNormalFormatter('platform_title')
        },
        {
            label: $t('adPlatformAccount.itemFields.data'),
            prop: 'data',
            minWidth: 100,
            renderHeader,
            formatter: createNormalFormatter('data')
        },
        {
            label: $t('adPlatformAccount.itemFields.status'),
            prop: 'status',
            minWidth: 130,
            renderHeader,
            cellRenderer: statusRender
        },
        {
            label: $t('adPlatformAccount.itemFields.create_time'),
            prop: 'create_time',
            minWidth: 180,
            renderHeader,
            cellRenderer: createTimestampRender('create_time')
        },
        {
            label: $t('adPlatformAccount.itemFields.update_time'),
            prop: 'update_time',
            minWidth: 180,
            renderHeader,
            cellRenderer: createTimestampRender('update_time')
        }, 
        {
            label: $t('common.actions'),
            width: 160,
            slot: 'operation',
            fixed: 'right',
            renderHeader
        }
    ];

    const adPlatformAccountPageOptsStore = usePageOptionsStoreHook();

    const controller = new AbortController();

    const loading = ref(false);
    const list = ref<Array<AdPlatformAccountItem>>([]);

    const pagination = reactive<PaginationProps>({
        total: 0,
        pageSize: adPlatformAccountPageOptsStore.pageSize,
        currentPage: adPlatformAccountPageOptsStore.currentPage,
        background: true,
        pageSizes: [5, 10, 15, 20, 50, 100, 10000, 100000]
    });

    const sortBy = ref(adPlatformAccountPageOptsStore.sortBy);

    const searchOptions = computed(() => adPlatformAccountPageOptsStore.searchOpts);

    const convertSortBy = () => {
        if (sortBy.value?.prop && sortBy.value?.order) {
            return {
                sort_by: [{ column: sortBy.value.prop, asc: String(sortBy.value.order).indexOf('desc') == -1 }]
            };
        }
        return {};
    };

    const loadList = () => {
        if (loading.value) return;

        loading.value = true;

        const data: ListReq = {
            page_no: pagination.currentPage,
            page_size: pagination.pageSize,
            ...convertSortBy(),
            ...searchOptions2ListReq(searchOptions.value)
        };

        fetchList({ data, signal: controller.signal }).then(
            res => {
                loading.value = false;
                list.value = res?.data?.list ?? [];
                pagination.total = res?.data?.total ?? 0;
                adPlatformAccountPageOptsStore.updateCurrentPage(pagination.currentPage);
                adPlatformAccountPageOptsStore.updatePageSize(pagination.pageSize);
                adPlatformAccountPageOptsStore.updateSortBy(sortBy.value);
                adPlatformAccountPageOptsStore.recordState();
            },
            err => {
                loading.value = false;
                if (!err?.isCancelRequest) {
                    message(err, { type: 'error' });
                }
            }
        );
    };

    const onRefreshData = () => loadList();

    const vLoading = resolveDirective('loading');

    const refCreatorForm = ref<AdPlatformAccountEditorFormInstance>(null);
    const refCreatorFormMode = ref<string>('dialog');
    const defaultAdPlatformAccountItem = (): AdPlatformAccountItem => ({ status: CommonStatusActivate });
    const creatorFormData = ref<AdPlatformAccountItem>(defaultAdPlatformAccountItem());
    const creatorSubmitting = ref<boolean>(false);
    const onCreateAdPlatformAccount = () => {
        addDialog({
            draggable: true,
            normalClass: 'no-padding-body-dlg !w-[95%] !max-w-[550px]',
            fullscreenClass: 'no-padding-body-dlg',
            ...CustomDialogHF(t('adPlatformAccount.createAdPlatformAccount'), { hasFullScreen: true }),
            fullScreenCallBack: ({ fullscreen }) => (refCreatorFormMode.value = fullscreen ? 'page' : 'dialog'),
            contentRenderer: () =>
                withDirectives(
                    createVNode(AdPlatformAccountEditorForm, {
                        ref: refCreatorForm,
                        isEdit: false,
                        mode: refCreatorFormMode.value,
                        adPlatformAccount: creatorFormData.value
                    }),
                    [[vLoading, creatorSubmitting.value]]
                ),
            beforeSure: done => {
                refCreatorForm.value?.validate((valid, _fields) => {
                    if (valid) {
                        if (creatorSubmitting.value) return;
                        doCreateAdPlatformAccount(clone(toRaw(creatorFormData.value)), done);
                    }
                });
            },
            closeCallBack: () => {
                creatorFormData.value = defaultAdPlatformAccountItem();
            }
        });
    };
    const doCreateAdPlatformAccount = (adPlatformAccount: AdPlatformAccountItem, done: Function) => {
        creatorSubmitting.value = true;
        const data = {
            id: adPlatformAccount.id,
            code: adPlatformAccount.code,
            name: adPlatformAccount.name,
            platform_code: adPlatformAccount.platform_code,
            platform_title: adPlatformAccount.platform_title,
            data: adPlatformAccount.data,
            status: adPlatformAccount.status,
            create_time: adPlatformAccount.create_time,
            update_time: adPlatformAccount.update_time,
        };
        fetchCreateAdPlatformAccount({ data }).then(
            () => {
                creatorSubmitting.value = false;
                message(
                    t('common.changeStatusMessage', {
                        action: t('common.create'),
                        content: adPlatformAccount.name
                    }),
                    { type: 'success' }
                );
                loadList();
                done();
            },
            err => {
                creatorSubmitting.value = false;
                message(
                    t('common.changeStatusFailed', {
                        action: t('common.create'),
                        content: adPlatformAccount.name,
                        msg: err
                    }),
                    { type: 'error' }
                );
            }
        );
    };

    const router = useRouter();
    const onEditAdPlatformAccount = (adPlatformAccount: AdPlatformAccountItem) => {
        const parameter: RouteParamsRaw = { id: `${adPlatformAccount.id}` };
        useMultiTagsStoreHook().handleTags('push', {
            path: '/config/adplatformaccount/item/:id',
            name: 'AdPlatformAccountEditor',
            params: parameter,
            meta: {
                title: t('adPlatformAccount.editAdPlatformAccountById', { id: adPlatformAccount.id }),
                keepAlive: true
            }
        });
        router.push({ name: 'AdPlatformAccountEditor', params: parameter });
    };

    const { onDeleteItem: onDeleteAdPlatformAccount } = useDeleteItemHook<AdPlatformAccountItem, DeleteReq>({
        fetchDelete: fetchDeleteAdPlatformAccount,
        getDeleteParams: item => ({ ids: [item.id] }),
        getKey: item => item.id,
        getName: item => item.name,
        callback: () => loadList()
    });

    const onChangePageSize = (val: number) => {
        pagination.pageSize = val;
        loadList();
    };

    const onChangeCurrentPage = (val: number) => {
        pagination.currentPage = val;
        loadList();
    };

    const onSortChange = ({ prop, order }: Sort) => {
        sortBy.value = { prop, order };
        loadList();
    };

    onMounted(() => loadList());

    onUnmounted(() => {
        if (loading.value) {
            controller.abort();
            loading.value = false;
        }
    });

    return {
        columns,
        loading,
        pagination,
        sortBy,
        list,
        onRefreshData,
        onCreateAdPlatformAccount,
        onEditAdPlatformAccount,
        onDeleteAdPlatformAccount,
        onChangePageSize,
        onChangeCurrentPage,
        onSortChange
    };
}
