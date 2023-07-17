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
            <el-form-item v-if="isEdit" :label="$t('configMgr.itemFields.id')">
                <el-input v-model="formData.id" disabled />
            </el-form-item>
            <el-form-item :label="$t('configMgr.itemFields.name')" prop="name">
                <el-input v-model="formData.name" />
            </el-form-item>
            <el-form-item :label="$t('configMgr.itemFields.code')" prop="code">
                <el-input v-model="formData.code" />
            </el-form-item>
            <el-form-item :label="$t('configMgr.configEditor.typeId')" prop="type_id">
                <el-select-v2
                    class="!w-full"
                    v-model="formData.type_id"
                    :options="schemaOptions"
                    @change="onChangeSchemaId"
                />
            </el-form-item>
            <el-form-item :label="$t('configMgr.itemFields.type_code')" prop="type_code">
                <el-input v-model="formData.type_code" disabled />
            </el-form-item>
            <el-form-item :label="$t('configMgr.itemFields.schema')" prop="schema">
                <el-input
                    v-model="formData.schema"
                    type="textarea"
                    readonly
                    resize="none"
                    @click="onClickedSchemaInput"
                />
            </el-form-item>
            <el-form-item :label="$t('configMgr.itemFields.data')" prop="data">
                <el-input v-model="formData.data" type="textarea" readonly resize="none" @click="onClickedDataInput" />
            </el-form-item>
            <el-form-item :label="$t('configMgr.itemFields.remark')" prop="remark">
                <el-input v-model="formData.remark" type="textarea" resize="none" />
            </el-form-item>
            <el-form-item :label="$t('appManager.appEditor.status')" prop="status">
                <common-status-selector class="!w-full" v-model="formData.status" @change="onChangeStatus" />
            </el-form-item>
        </el-form>
        <el-button v-if="mode == 'page'" class="w-full" type="primary" @click="onClickedSubmit">{{
            $t('common.submit')
        }}</el-button>
    </el-scrollbar>
</template>

<script setup lang="ts">
    import { ref } from 'vue';
    import { type FormInstance } from 'element-plus';
    import { CommonStatusSelector } from '@/common/CommonStatusSelector';

    import { useForm, type ConfigEditorFormProps } from './formHooks';
    import { useRules } from './formRules';

    import { CommonStatusActivate } from '@/defines';
    import { type ConfigItem } from '@/api/config';

    defineOptions({ name: 'ConfigEditorForm' });

    const props = withDefaults(defineProps<ConfigEditorFormProps>(), {
        mode: 'dialog',
        isEdit: false,
        config: (): ConfigItem => ({ status: CommonStatusActivate })
    });

    const emit = defineEmits<{
        (evt: 'submit'): void;
    }>();

    const refForm = ref<FormInstance>();

    const {
        formData,
        schemaOptions,
        dataErrors,
        usedSchema,
        onChangeSchemaId,
        onClickedSchemaInput,
        onClickedDataInput,
        onChangeStatus,
        onClickedSubmit,
        validate,
        validateField
    } = useForm(refForm, props, emit);

    const formRules = useRules(usedSchema, dataErrors).formRules;

    defineExpose({ formData, validate, validateField });
</script>

<style lang="scss"></style>