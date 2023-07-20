<template>
    <el-scrollbar class="!px-[20px]" :max-height="mode == 'dialog' ? '300px' : undefined">
        <el-form
            ref="refForm"
            class="!my-[30px]"
            :model="formData"
            :rules="formRules"
            label-width="100px"
            size="small"
            label-position="left"
        >
            <el-form-item :label="$t('adPlatformAccount.itemFields.id')" prop="id">
                <el-input v-model="formData.id" />
            </el-form-item>

            <el-form-item :label="$t('adPlatformAccount.itemFields.code')" prop="code">
                <el-input v-model="formData.code" />
            </el-form-item>

            <el-form-item :label="$t('adPlatformAccount.itemFields.name')" prop="name">
                <el-input v-model="formData.name" />
            </el-form-item>

            <el-form-item :label="$t('adPlatformAccount.itemFields.platform_code')" prop="platform_code">
                <el-input v-model="formData.platform_code" />
            </el-form-item>

            <el-form-item :label="$t('adPlatformAccount.itemFields.platform_title')" prop="platform_title">
                <el-input v-model="formData.platform_title" />
            </el-form-item>

            <el-form-item :label="$t('adPlatformAccount.itemFields.data')" prop="data">
                <el-input v-model="formData.data" />
            </el-form-item>

            <el-form-item :label="$t('common.status')" prop="status">
                <common-status-selector class="!w-full" v-model="formData.status" @change="onChangeStatus" />
            </el-form-item>
            

            <el-form-item :label="$t('adPlatformAccount.itemFields.create_time')" prop="create_time">
                <el-input v-model="formData.create_time" />
            </el-form-item>

            <el-form-item :label="$t('adPlatformAccount.itemFields.update_time')" prop="update_time">
                <el-input v-model="formData.update_time" />
            </el-form-item>
            
        </el-form>
        <el-button v-if="mode == 'page'" class="w-full" type="primary" @click="onClickedSubmit">$t('common.submit')</el-button>
    </el-scrollbar>
</template>

<script setup lang="ts">
    import { ref } from 'vue';
    import { type FormInstance } from 'element-plus';
    import { CommonStatusSelector } from '@/common/CommonStatusSelector';

    import { useForm, type AdPlatformAccountEditorFormProps } from './formHooks';
    import { useRules } from './formRules';

    import { CommonStatusActivate } from '@/defines';
    import { type AdPlatformAccountItem } from '@/api/ad-platform-account';

    defineOptions({ name: 'AdPlatformAccountEditorForm' });

    const props = withDefaults(defineProps<AdPlatformAccountEditorFormProps>(), {
        mode: 'dialog',
        isEdit: false,
        config: (): AdPlatformAccountItem => ({ status: CommonStatusActivate })
    });

    const emit = defineEmits<{
        (evt: 'submit'): void;
    }>();

    const refForm = ref<FormInstance>();

    const {
        formData,
        dataErrors,
        onChangeSchemaId,
        onClickedSchemaInput,
        onClickedDataInput,
        onChangeStatus,
        onClickedSubmit,
        validate,
        validateField
    } = useForm(refForm, props, emit);

    const formRules = useRules(dataErrors).formRules;

    defineExpose({ formData, validate, validateField });
</script>

<style lang="scss"></style>