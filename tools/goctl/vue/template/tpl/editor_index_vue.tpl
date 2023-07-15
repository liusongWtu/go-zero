<template>
    <config-editor-form
        v-loading="loading"
        ref="refEditorForm"
        mode="page"
        :is-edit="true"
        :config="configItem"
        @submit="onSubmit"
    />
</template>

<script setup lang="ts">
    import { ref, computed, onMounted, onUnmounted } from 'vue';
    import { useRoute } from 'vue-router';
    import { useI18n } from 'vue-i18n';

    import { ConfigEditorForm, type ConfigEditorFormInstance } from './form';

    import { fetchConfigItem, fetchUpdateConfig, type ConfigItem, type ItemReq } from '@/api/config';

    import { message } from '@/utils/message';

    defineOptions({ name: 'ConfigEditorPage' });

    const refEditorForm = ref<ConfigEditorFormInstance>(null);

    const { t } = useI18n();
    const route = useRoute();

    const configId = computed(() => Number(route.params.id));

    const loading = ref<boolean>(false);
    const configItem = ref<ConfigItem>({});

    const controller = new AbortController();

    const loadConfigItemData = (id: number) => {
        if (loading.value) controller.abort();

        loading.value = true;
        const data: ItemReq = { id };
        fetchConfigItem({ data, signal: controller.signal })
            .then(
                res => {
                    configItem.value = res.data;
                },
                err => {
                    if (!err?.isCancelRequest) {
                        message(t('common.dataLoadFailed', { error: String(err) }), { type: 'error' });
                    }
                }
            )
            .finally(() => {
                loading.value = false;
            });
    };

    onMounted(() => loadConfigItemData(configId.value));

    onUnmounted(() => {
        if (loading.value) {
            controller.abort();
            loading.value = false;
        }
    });

    const onSubmit = () => {
        refEditorForm.value?.validate((valid, _fields) => {
            if (valid) {
                doUpdateConfig(refEditorForm.value.formData);
            }
        });
    };

    const doUpdateConfig = (config: ConfigItem) => {
        const data = {
            id: configId.value,
            name: config.name,
            data: config.data,
            code: config.code,
            type_id: config.type_id,
            type_name: config.type_name,
            type_code: config.type_code,
            schema: config.schema,
            remark: config.remark,
            status: config.status
        };
        fetchUpdateConfig({ data }).then(
            () => {
                message(t('common.changeStatusMessage', { action: t('common.update'), content: config.name }), {
                    type: 'success'
                });
            },
            err => {
                message(
                    t('common.changeStatusFailed', {
                        action: t('common.update'),
                        content: config.name,
                        msg: err
                    }),
                    { type: 'error' }
                );
            }
        );
    };
</script>

<style lang="scss" scoped></style>
