<template>
    <search-pad @reset="onReset" @search="doDebounceSearch">
        <el-form ref="refForm" size="small" :inline="true" :model="formData" :rules="formRules" class="px-2">
            <responsive-row>
                <responsive-col>
                    <el-form-item :label="$t('adPlatformAccount.search.dateRange')">
                        <date-range-picker v-model="formData.date" @change="onSearch" />
                    </el-form-item>
                </responsive-col>
                <responsive-col>
                    <el-form-item :label="$t('adPlatformAccount.itemFields.status')">
                        <common-status-selector
                            class="!w-full"
                            v-model="formData.status"
                            multiple
                            clearable
                            @change="doDebounceSearch"
                        />
                    </el-form-item>
                </responsive-col>
                <responsive-col>
                    <el-form-item :label="$t('adPlatformAccount.itemFields.id')" prop="ids">
                        <el-input
                            v-model="formData.ids"
                            class="!w-full"
                            :placeholder="$t('adPlatformAccount.search.placeholderId')"
                            clearable
                            @input="doDebounceSearch"
                        />
                    </el-form-item>
                </responsive-col>
                <responsive-col>
                    <el-form-item :label="$t('adPlatformAccount.itemFields.name')">
                        <el-input
                            v-model="formData.name"
                            :placeholder="$t('adPlatformAccount.search.placeholderName')"
                            clearable
                            @input="doDebounceSearch"
                        />
                    </el-form-item>
                </responsive-col>
                <responsive-col>
                    <el-form-item :label="$t('adPlatformAccount.itemFields.code')">
                        <el-input
                            v-model="formData.code"
                            :placeholder="$t('adPlatformAccount.search.placeholderCode')"
                            clearable
                            @input="doDebounceSearch"
                        />
                    </el-form-item>
                </responsive-col>
            </responsive-row>
        </el-form>
    </search-pad>
</template>

<script setup lang="ts">
    import { ref } from 'vue';
    import { FormInstance } from 'element-plus';

    import { SearchPad } from '@/common/SearchPad';
    import { ResponsiveRow, ResponsiveCol } from '@/common/ResponsiveGrid';

    import { DateRangePicker } from '@/common/DateRangePicker';
    import { CommonStatusSelector } from '@/common/CommonStatusSelector';

    import { useForm } from './formHook';
    import { useRules } from './formRules';


    const emit = defineEmits<{
        (evt: 'search'): void;
    }>();

    const refForm = ref<FormInstance>();

    const { formData, onReset, onSearch, doDebounceSearch } = useForm(refForm, emit);

    const formRules = useRules().formRules;
</script>

<style lang="scss" scoped></style>