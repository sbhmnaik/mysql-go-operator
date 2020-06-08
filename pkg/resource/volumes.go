package resource

import (
	"github.com/persistentsys/mysql-go-operator/pkg/apis/mysql/v1alpha1"
	"github.com/persistentsys/mysql-go-operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var volLog = logf.Log.WithName("resource_volumes")

// GetMysqlVolumeName - return name of PV used in MySQL
func GetMysqlVolumeName(v *v1alpha1.MySQL) string {
	return v.Name + "-" + v.Namespace + "-pv"
}

// GetMysqlVolumeClaimName - return name of PVC used in MySQL
func GetMysqlVolumeClaimName(v *v1alpha1.MySQL) string {
	return v.Name + "-pv-claim"
}

// GetMysqlBkpVolumeName - return name of PV used in DB Backup
// func GetMysqlBkpVolumeName(bkp *v1alpha1.Backup) string {
// 	return bkp.Name + "-" + bkp.Namespace + "-pv"
// }

// GetMysqlBkpVolumeClaimName - return name of PVC used in DB Backup
// func GetMysqlBkpVolumeClaimName(bkp *v1alpha1.Backup) string {
// 	return bkp.Name + "-pv-claim"
// }

// NewDbBackupPV Create a new PV object for Database Backup
// func NewDbBackupPV(bkp *v1alpha1.Backup, v *v1alpha1.MySQL, scheme *runtime.Scheme) *corev1.PersistentVolume {
// 	volLog.Info("Creating new PV for Database Backup")
// 	labels := utils.MySQLBkpLabels(bkp, "mysql-backup")
// 	pv := &corev1.PersistentVolume{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: GetMysqlBkpVolumeName(bkp),
// 			// Namespace: v.Namespace,
// 			Labels: labels,
// 		},
// 		Spec: corev1.PersistentVolumeSpec{
// 			StorageClassName: "manual",
// 			Capacity: corev1.ResourceList{
// 				corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(bkp.Spec.BackupSize),
// 			},
// 			AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
// 			PersistentVolumeSource: corev1.PersistentVolumeSource{
// 				HostPath: &corev1.HostPathVolumeSource{
// 					Path: bkp.Spec.BackupPath},
// 			},
// 		},
// 	}

// 	volLog.Info("PV created for Database Backup ")
// 	controllerutil.SetControllerReference(bkp, pv, scheme)
// 	return pv
// }

// NewDbBackupPVC Create a new PV Claim object for Database Backup
// func NewDbBackupPVC(bkp *v1alpha1.Backup, v *v1alpha1.MySQL, scheme *runtime.Scheme) *corev1.PersistentVolumeClaim {
// 	volLog.Info("Creating new PVC for Database Backup")
// 	labels := utils.MySQLBkpLabels(bkp, "mysql-backup")
// 	storageClassName := "manual"
// 	pvc := &corev1.PersistentVolumeClaim{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      GetMysqlBkpVolumeClaimName(bkp),
// 			Namespace: v.Namespace,
// 			Labels:    labels,
// 		},
// 		Spec: corev1.PersistentVolumeClaimSpec{
// 			StorageClassName: &storageClassName,
// 			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
// 			Resources: corev1.ResourceRequirements{
// 				Requests: corev1.ResourceList{
// 					corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(bkp.Spec.BackupSize),
// 				},
// 			},
// 			VolumeName: GetMysqlBkpVolumeName(bkp),
// 		},
// 	}

// 	volLog.Info("PVC created for Database Backup ")
// 	controllerutil.SetControllerReference(bkp, pvc, scheme)
// 	return pvc
// }

// NewMySqlPV Create a new PV object for MySQL
func NewMySqlPV(v *v1alpha1.MySQL, scheme *runtime.Scheme) *corev1.PersistentVolume {
	volLog.Info("Creating new PV for MySQL")
	labels := utils.Labels(v, "mysql")
	pv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: GetMysqlVolumeName(v),
			// Namespace: v.Namespace,
			Labels: labels,
		},
		Spec: corev1.PersistentVolumeSpec{
			StorageClassName: "manual",
			Capacity: corev1.ResourceList{
				corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(v.Spec.DataStorageSize),
			},
			AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: v.Spec.DataStoragePath},
			},
		},
	}

	volLog.Info("PV created for MySQL ")
	controllerutil.SetControllerReference(v, pv, scheme)
	return pv
}

// NewMySqlPVC Create a new PV Claim object for MySQL
func NewMySqlPVC(v *v1alpha1.MySQL, scheme *runtime.Scheme) *corev1.PersistentVolumeClaim {
	volLog.Info("Creating new PVC for MySQL")
	labels := utils.Labels(v, "mysql")
	storageClassName := "manual"
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      GetMysqlVolumeClaimName(v),
			Namespace: v.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			StorageClassName: &storageClassName,
			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(v.Spec.DataStorageSize),
				},
			},
			VolumeName: GetMysqlVolumeName(v),
		},
	}

	volLog.Info("PVC created for MySQL ")
	controllerutil.SetControllerReference(v, pvc, scheme)
	return pvc
}
