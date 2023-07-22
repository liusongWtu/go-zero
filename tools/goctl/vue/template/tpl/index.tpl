<template>
    <div class="{{.kebabObject}}-manager-page">
        <search @search="onRefreshData" />
        <pure-table-bar :title="{{ printf "$t('%s.%sList')" .lowerStartCamelObject .lowerStartCamelObject}}" :columns="columns" @refresh="onRefreshData">
            <template #buttons>
                <el-button type="primary" :icon="useRenderIcon(AddFill)" @click="onCreate{{.upperStartCamelObject}}">
                    {{"{{"}} $t('{{.lowerStartCamelObject}}.create{{.upperStartCamelObject}}')  {{"}}"}}
                </el-button>
            </template>
            <template v-slot="{ size, dynamicColumns }">
                <pure-table
                    border
                    align-whole="center"
                    table-layout="auto"
                    :size="size"
                    :loading="loading"
                    :data="list"
                    :columns="dynamicColumns"
                    :pagination="pagination"
                    :default-sort="sortBy"
                    :paginationSmall="size === 'small' ? true : false"
                    :header-cell-style="{
                        background: 'var(--el-table-row-hover-bg-color)',
                        color: 'var(--el-text-color-primary)'
                    }"
                    @page-size-change="onChangePageSize"
                    @page-current-change="onChangeCurrentPage"
                    @sort-change="onSortChange"
                >
                    <template #operation="{ row }">
                        <el-button
                            class="reset-margin"
                            link
                            type="primary"
                            :size="size"
                            :icon="useRenderIcon(EditPen)"
                            @click="onEdit{{.upperStartCamelObject}}(row)"
                        >
                             {{"{{"}} $t('common.edit')  {{"}}"}}
                        </el-button>

                        <el-popconfirm
                            :title="$t('common.deleteItemConfirm', { content: row.name })"
                            @confirm="onDelete{{.upperStartCamelObject}}(row)"
                        >
                            <template #reference>
                                <el-button
                                    class="reset-margin"
                                    link
                                    type="primary"
                                    :size="size"
                                    :icon="useRenderIcon(Delete)"
                                >
                                     {{"{{"}} $t('common.delete')  {{"}}"}}
                                </el-button>
                            </template>
                        </el-popconfirm>
                    </template>
                </pure-table>
            </template>
        </pure-table-bar>
    </div>
</template>

<script setup lang="ts">
    import Search from './search/form.vue';
    import { PureTableBar } from '@/components/RePureTableBar';

    import { useRenderIcon } from '@/components/ReIcon/src/hooks';

    import Delete from '@iconify-icons/ep/delete';
    import EditPen from '@iconify-icons/ep/edit-pen';
    import AddFill from '@iconify-icons/ri/add-circle-line';

    import { useList } from './listHook';

    defineOptions({ name: '{{.upperStartCamelObject}}Manager' });

    const {
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
    } = useList();
</script>

<style lang="scss" scoped>
    .{{.kebabObject}}-manager-page {
        :deep(.el-table__row .cell) {
            margin: auto;
            text-align: center;
        }
    }
</style>