import { ref, reactive, computed, toRaw, createVNode, onMounted, onUnmounted } from 'vue';
import { useRouter, type RouteParamsRaw } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { type Sort } from 'element-plus';
import { type PaginationProps } from '@pureadmin/table';

import { JsonCodeMirror } from '@/components/JsonCodeMirror';
import { addDialog } from '@/components/ReDialog';
import { CustomDialogHF } from '@/common/CustomDialogHF';
import { {{.upperStartCamelObject}}EditorForm, type {{.upperStartCamelObject}}EditorFormInstance } from './editor/form';

import {
    createNormalFormatter,
    noWarpContent,
    createTimestampRender,
    createCommonStatusRender
} from '@/common/tableCellRenders';

import clone from 'ramda/es/clone';

import {
    fetchCreate{{.upperStartCamelObject}},
    fetchUpdate{{.upperStartCamelObject}},
    fetchDelete{{.upperStartCamelObject}},
    fetchList,
    type {{.upperStartCamelObject}}Item,
    type ListReq,
    type DeleteReq
} from '@/api/{{.kebabObject}}';
import { usePageOptionsStoreHook, searchOptions2ListReq } from '@/store/modules/{{.lowerStartCamelObject}}PageOpts';

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
        fetchUpdate: fetchUpdate{{.upperStartCamelObject}},
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
        {{.tableFields}} 
        {
            label: $t('common.actions'),
            width: 160,
            slot: 'operation',
            fixed: 'right',
            renderHeader
        }
    ];

    const {{.lowerStartCamelObject}}PageOptsStore = usePageOptionsStoreHook();

    const controller = new AbortController();

    const loading = ref(false);
    const list = ref<Array<{{.upperStartCamelObject}}Item>>([]);

    const pagination = reactive<PaginationProps>({
        total: 0,
        pageSize: {{.lowerStartCamelObject}}PageOptsStore.pageSize,
        currentPage: {{.lowerStartCamelObject}}PageOptsStore.currentPage,
        background: true,
        pageSizes: [5, 10, 15, 20, 50, 100, 10000, 100000]
    });

    const sortBy = ref({{.lowerStartCamelObject}}PageOptsStore.sortBy);

    const searchOptions = computed(() => {{.lowerStartCamelObject}}PageOptsStore.searchOpts);

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
                {{.lowerStartCamelObject}}PageOptsStore.updateCurrentPage(pagination.currentPage);
                {{.lowerStartCamelObject}}PageOptsStore.updatePageSize(pagination.pageSize);
                {{.lowerStartCamelObject}}PageOptsStore.updateSortBy(sortBy.value);
                {{.lowerStartCamelObject}}PageOptsStore.recordState();
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

    const refCreatorForm = ref<{{.upperStartCamelObject}}EditorFormInstance>(null);
    const default{{.upperStartCamelObject}}Item = (): {{.upperStartCamelObject}}Item => ({ status: CommonStatusActivate });
    const formDataCreate = ref<{{.upperStartCamelObject}}Item>(default{{.upperStartCamelObject}}Item());
    const onCreate{{.upperStartCamelObject}} = () => {
        addDialog({
            draggable: true,
            ...CustomDialogHF(t('{{.lowerStartCamelObject}}.create{{.upperStartCamelObject}}'), 'no-padding-body-dlg !w-[95%] !max-w-[550px]'),
            contentRenderer: () =>
                createVNode({{.upperStartCamelObject}}EditorForm, {
                    ref: refCreatorForm,
                    isEdit: false,
                    mode: 'dialog',
                    schema: formDataCreate
                }),
            beforeSure: done => {
                refCreatorForm.value?.validate((valid, _fields) => {
                    if (valid) {
                        const data = clone(toRaw(formDataCreate.value));
                        formDataCreate.value = default{{.upperStartCamelObject}}Item();
                        doCreate{{.upperStartCamelObject}}(data, done);
                    }
                });
            },
            closeCallBack: () => {
                formDataCreate.value = default{{.upperStartCamelObject}}Item();
            }
        });
    };
    const doCreate{{.upperStartCamelObject}} = ({{.lowerStartCamelObject}}: {{.upperStartCamelObject}}Item, done: Function) => {
        const data = {
{{.requestField}}
        };
        fetchCreate{{.upperStartCamelObject}}({ data }).then(
            () => {
                message(
                    t('common.changeStatusMessage', {
                        action: t('common.create'),
                        content: {{.lowerStartCamelObject}}.name
                    }),
                    { type: 'success' }
                );
                loadList();
                done();
            },
            err => {
                message(
                    t('common.changeStatusFailed', {
                        action: t('common.create'),
                        content: {{.lowerStartCamelObject}}.name,
                        msg: err
                    }),
                    { type: 'error' }
                );
            }
        );
    };

    const router = useRouter();
    const onEdit{{.upperStartCamelObject}} = ({{.lowerStartCamelObject}}: {{.upperStartCamelObject}}Item) => {
        const parameter: RouteParamsRaw = { id: `${{{.lowerStartCamelObject}}.id}` };
        useMultiTagsStoreHook().handleTags('push', {
            path: '/config/schema/item/:id',
            name: '{{.upperStartCamelObject}}Editor',
            params: parameter,
            meta: {
                title: t('{{.lowerStartCamelObject}}.edit{{.upperStartCamelObject}}ById', { id: {{.lowerStartCamelObject}}.id }),
                keepAlive: true
            }
        });
        router.push({ name: '{{.upperStartCamelObject}}Editor', params: parameter });
    };

    const { onDeleteItem: onDelete{{.upperStartCamelObject}} } = useDeleteItemHook<{{.upperStartCamelObject}}Item, DeleteReq>({
        fetchDelete: fetchDelete{{.upperStartCamelObject}},
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
        onCreate{{.upperStartCamelObject}},
        onEdit{{.upperStartCamelObject}},
        onDelete{{.upperStartCamelObject}},
        onChangePageSize,
        onChangeCurrentPage,
        onSortChange
    };
}
