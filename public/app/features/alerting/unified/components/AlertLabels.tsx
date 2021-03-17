import React, { useState } from 'react';

import { Button, Field, Input, Label, stylesFactory } from '@grafana/ui';
import { GrafanaTheme } from '@grafana/data';
import { config } from 'app/core/config';
import { css } from 'emotion';

interface Props {}

type PromLabel = {
  [key: string]: string;
};

const AlertLabels = (props: Props) => {
  const { labels, addLabel } = useLabels();
  const labelName = useInput();
  const labelValue = useInput();

  const styles = getStyles(config.theme);

  const labelPreview = () => {
    return (
      <ul>
        {Object.entries(labels).map(([labelKey, labelValue]) => {
          return <li key={`${labelKey}-${labelValue}`}>{`${labelKey}=${labelValue}`}</li>;
        })}
      </ul>
    );
  };

  const clearFields = () => {
    labelName.clearValue();
    labelValue.clearValue();
  };

  return (
    <>
      <Label>Custom labels</Label>
      {labels && labelPreview()}
      <div className={styles.flexRow}>
        <Field className={styles.formInput}>
          <Input placeholder="label" value={labelName.value} onChange={labelName.handleChange} />
        </Field>
        <Field className={styles.formInput}>
          <Input placeholder="value" value={labelValue.value} onChange={labelValue.handleChange} />
        </Field>
      </div>
      <Button
        disabled={!labelName.value || !labelValue.value}
        type="button"
        variant="secondary"
        size="sm"
        onClick={() => {
          if (labelName.value && labelValue.value) {
            addLabel({ [labelName.value]: labelValue.value });
          }
          clearFields();
        }}
      >
        Add label
      </Button>
    </>
  );
};

const useInput = () => {
  const [value, setValue] = useState<string>();

  const handleChange = (event: any) => {
    event.preventDefault();
    setValue(event.target.value);
  };

  const clearValue = () => {
    setValue('');
  };

  return { value, handleChange, clearValue };
};

const useLabels = () => {
  const [labels, setLabels] = useState<PromLabel>({});

  const addLabel = (label: PromLabel) => {
    setLabels({ ...labels, ...label });
  };

  return { labels, addLabel };
};

const getStyles = stylesFactory((theme: GrafanaTheme) => {
  return {
    formInput: css`
      width: 400px;
      & + & {
        margin-left: ${theme.spacing.sm};
      }
    `,
    flexRow: css`
      display: flex;
      flex-direction: row;
      justify-content: flex-start;
    `,
  };
});

export default AlertLabels;
